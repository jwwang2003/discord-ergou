package structs

import (
	"errors"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"github.com/kkdai/youtube/v2"
)

type VoiceInstance struct {
	VoiceConn 	*discordgo.VoiceConnection
	Session 		*discordgo.Session
	encoder 		*dca.EncodeSession
	stream 			*dca.StreamingSession
	queueMutex 	sync.Mutex
	audioMutex 	sync.Mutex
	nowPlaying 	*Song
	queue 			[]Song
	GuildID 		string
	ChannelID 	string
	speaking 		bool
	pause 			bool
	stop 				bool
	skip 				bool
}

// basic methods for controlling the voice instance

func (v *VoiceInstance) PlayQueue(song Song) ( bool ) {
	v.QueuePrepend(song)
	if v.speaking {
		return true
	}

	go func() {
		v.audioMutex.Lock()
		defer v.audioMutex.Unlock()

		for {
			if len(v.queue) == 0 {
				log.Println("The queue is empty!")
				return
			}

			v.nowPlaying = v.QueueGet()

			log.Println("Now playing: ", v.nowPlaying.Title)

			v.stop = false
			v.skip = false
			v.speaking = true
			v.pause = false
			v.VoiceConn.Speaking(true)

			v.DCA(v.nowPlaying.VideoURL)

			// if(v.stop) {
			// 	v.QueueClean()
			// }

			v.stop = false
			v.skip = false
			v.speaking = false
			v.VoiceConn.Speaking(false)
		}
	} ()
	return true
}

func (v *VoiceInstance) DCA(url string) {
	// real-time encoding configuration
	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 128
	options.Application = "audio"	// favors quality over delay

	var videoID string = "LL-gyhZVvx0"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		log.Print("Error getting video data")
	}

	formats := video.Formats.WithAudioChannels()
	downloadURL, err := client.GetStreamURL(video, &formats[0])
	if err != nil {
		log.Print("Error getting downloadURL")
	}

	encodeSession, err := dca.EncodeFile(downloadURL, options)
	if err !=  nil {
		log.Println("Failed to create an encoding session: ", err)
	}

	v.encoder = encodeSession
	done := make(chan error)
	stream := dca.NewStream(encodeSession, v.VoiceConn, done)
	v.stream = stream

	select {
	case err := <-done:
		if err != nil && err != io.EOF {
			log.Println("An error occured while encoding and streaming audio: ", err)
		}
		fmt.Print("Checkpoint")
		encodeSession.Cleanup()
		return
	}
}

func (v *VoiceInstance) Skip() ( bool, error ) {
	if v.speaking {
		if v.encoder != nil {
			v.encoder.Cleanup()
			return true, nil
		}
	}
	return false, errors.New("failed to skip track")
}

func (v *VoiceInstance) Stop() ( bool, error ) {
	v.stop = true
	if v.encoder != nil {
		v.encoder.Cleanup()
		return true, nil
	}
	return false, errors.New("no audio is being encoded at the momment")
}

func (v *VoiceInstance) Pause() ( bool, error ) {
	v.pause = true
	if v.stream != nil {
		v.stream.SetPaused(true)
		return true, nil
	}
	return false, errors.New("no audio is being streamed at the momment")
}

func (v *VoiceInstance) Resume() ( bool, error ) {
	v.pause = false
	if v.stream != nil {
		v.stream.SetPaused(false)
		return true, nil
	}
	return false, errors.New("unable to resume, no stream present")
}


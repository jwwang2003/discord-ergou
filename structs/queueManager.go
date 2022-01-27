package structs

// FIFO - first in first out
// methods for interacting with queue in VoiceInstance

// add to end of queue
func (v *VoiceInstance) QueueAppend(song Song) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()

	v.queue = append(v.queue, song)
}

// add to the front of queue 
// (behind the first element if voice instance is active)
func (v *VoiceInstance) QueuePrepend(song Song) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()

	if v.nowPlaying != nil {
		temp := append([]Song{v.queue[0]}, []Song{song}...)
		v.queue = append(temp, v.queue[1:]...)
	} else {
		v.queue = append([]Song{song}, v.queue...)
	}
}

func (v *VoiceInstance) QueueGet() (*Song) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()

	if len(v.queue) != 0 {
		var song Song
		song, v.queue = v.queue[0], v.queue[1:]
		return &song
	}

	return nil
}

// clear everything from queue
func (v *VoiceInstance) QueueClean() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	v.queue = []Song{}
}
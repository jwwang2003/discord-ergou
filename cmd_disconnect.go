package main

import (
	. "ergou/structs"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

var Disconnect = discordgo.ApplicationCommand {
	Name: 				"disconnect",
	Description: 	"Disconnect from current voice channel",
	Type: 				discordgo.ChatApplicationCommand,
	Options: 			[]*discordgo.ApplicationCommandOption {},
}

func DisconnectHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var v *VoiceInstance = voiceInstances[i.GuildID]
		
		if v == nil {
			log.Println("Bot is not part of a voice channel")
			return
		}

		v.Stop()
		time.Sleep(200 * time.Millisecond)

		v.VoiceConn.Disconnect()

		mutex.Lock()
		delete(voiceInstances, i.GuildID)
		mutex.Unlock()

		log.Println("Left the voice channel and deleted voice instance")
}
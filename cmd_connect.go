package main

// import (
// 	. "ergou/structs"
// 	"log"

// 	"github.com/bwmarrin/discordgo"
// )

// var Connect = discordgo.ApplicationCommand {
// 	Name: 				"connect",
// 	Description: 	"Join current voice channel",
// 	Type: 				discordgo.ChatApplicationCommand,
// 	Options: 			[]*discordgo.ApplicationCommandOption {},
// }

// func ConnectHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 		voiceChannelID, err := FindUserVoiceChannelFromGuild(s, i.GuildID, i.Member.User.ID)
// 		if err != nil { log.Println("error: ", err) }

// 		var v *VoiceInstance

// 		if voiceInstances[i.GuildID] != nil {
// 			log.Println("Voice instance already exists")
// 		} else {
// 			mutex.Lock()
// 			v = new(VoiceInstance)
// 			voiceInstances[i.GuildID] = v
// 			v.GuildID = i.GuildID
// 			v.Session = s
// 			mutex.Unlock()
// 		}

// 		v.VoiceConn, err = s.ChannelVoiceJoin(i.GuildID, voiceChannelID.ChannelID, false, true)
// 		if err != nil {
// 			v.Stop()
// 			log.Println("Error joining voice channel")
// 			return
// 		}
// 		v.VoiceConn.Speaking(false)
// 		log.Println("New voice instance created & joined the voice channel")

// 		v.PlayQueue(Song{
// 			Title: "Test",
// 			VideoURL: "google.com",
// 		})

// 		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Content: "Bot has joined current voice channel",
// 			},
// 		})
// 		if err != nil {
// 			log.Println("An error occured when sending messsage", err)
// 		}
// }
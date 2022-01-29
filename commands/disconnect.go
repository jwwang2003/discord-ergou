package commands

import (
	"sync"
	"time"

	"ergou/helpers"
	"ergou/structs"

	"github.com/bwmarrin/discordgo"
)

var Disconnect = discordgo.ApplicationCommand {
	Name: 				"disconnect",
	Description: 	"Disconnect from current voice channel",
	Type: 				discordgo.ChatApplicationCommand,
	Options: 			[]*discordgo.ApplicationCommandOption {},
}

func DisconnectHandler(
	s 	*discordgo.Session,
	i 	*discordgo.InteractionCreate,
	m 	*sync.Mutex,
	vi  map[string]*structs.VoiceInstance,
) {
		var v *structs.VoiceInstance = vi[i.GuildID]
		
		if v == nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title: "Mr. Shiba is not part of any voice channels",
							Color: helpers.ERROR,
							Author: &discordgo.MessageEmbedAuthor{
								IconURL: "https://cdn.discordapp.com/emojis/936535532677251122.gif",
							},
						},
					},
				},
			})
			return
		}

		v.Stop()
		time.Sleep(200 * time.Millisecond)

		v.VoiceConn.Disconnect()

		m.Lock()
		delete(vi, i.GuildID)
		m.Unlock()

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title: "Mr. Shiba has left the voice channel",
						Color: helpers.INFO,
						Author: &discordgo.MessageEmbedAuthor{
							IconURL: "https://cdn.discordapp.com/emojis/936535532677251122.gif",
						},
					},
				},
			},
		})
}
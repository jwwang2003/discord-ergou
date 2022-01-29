package commands

import (
	"fmt"
	"sync"

	"ergou/helpers"
	"ergou/structs"

	"github.com/bwmarrin/discordgo"
)

var Connect = discordgo.ApplicationCommand {
	Name: 				"connect",
	Description: 	"Join current voice channel",
	Type: 				discordgo.ChatApplicationCommand,
	Options: 			[]*discordgo.ApplicationCommandOption{},
}

func ConnectHandler(
	s 	*discordgo.Session,
	i 	*discordgo.InteractionCreate,
	m 	*sync.Mutex,
	vi  map[string]*structs.VoiceInstance,
) {
	// find which voice channel the "user" is in
	voiceState, err := helpers.FindUserVoiceState(s, i.GuildID, i.Member.User.ID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title: fmt.Sprintf("I %v", err),
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

	var v *structs.VoiceInstance

	if vi[i.GuildID] != nil {
		// voice instance already exists for this channel
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("%v is already in a voice channel for this server", s.State.User.Username),
			},
		})
		return
	} else {
		m.Lock()
		v = new(structs.VoiceInstance)
		vi[i.GuildID] = v
		v.GuildID = i.GuildID
		v.Session = s
		m.Unlock()
	}

	v.VoiceConn, err = s.ChannelVoiceJoin(i.GuildID, voiceState.VS.ChannelID, false, true)
	if err != nil {
		v.Stop()
		return
	}
	v.VoiceConn.Speaking(false)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Mr. Shiba at your service",
					Description: fmt.Sprintf("I have joined `%v` and will be bound to <#%v>", voiceState.Name, i.ChannelID),
					Color: helpers.SUCCESS,
					Author: &discordgo.MessageEmbedAuthor{
						Name: " ",
						IconURL: "https://cdn.discordapp.com/emojis/936535532677251122.gif",
					},
				},
			},
		},
	})
}
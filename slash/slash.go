package slash

import (
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name: "play",
			Description: "Enqueue a local song",
		},
		{
			Name: "youtube",
			Description: "Enqueue a song via keywords or url from YouTube",
		},
		{
			Name: "spotify",
			Description: "Enqueue a song via keywords or url from Spotify",
		},
	}
	
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		"play": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		},
		"youtube": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		},
		"spotify": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		},
	}
)
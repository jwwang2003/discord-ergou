package main

import (
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		&Connect,
		&Disconnect,
	}

	CommandHandlers = map[string]func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		"connect": ConnectHandler,
		"disconnect": DisconnectHandler,
	}
)
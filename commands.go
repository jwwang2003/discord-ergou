package main

import (
	"sync"

	"ergou/commands"
	"ergou/structs"

	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		&commands.Connect,
	}

	CommandHandlers = map[string]func(
		s 	*discordgo.Session,
		i 	*discordgo.InteractionCreate,
		m 	*sync.Mutex,
		vi  map[string]*structs.VoiceInstance,
	) {
		"connect": commands.ConnectHandler,
	}
)
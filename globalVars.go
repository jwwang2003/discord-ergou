package main

import (
	. "ergou/structs"
	"sync"

	"github.com/bwmarrin/discordgo"
)

var (
	discordSession 	*discordgo.Session
	// keep track of the voice connection state to each guild
	voiceInstances 	= map[string]*VoiceInstance{}
	// syncronizing goroutines to prevent race conditions
	mutex 					sync.Mutex
	// song queueing channel
	songChan 				chan SongPkg
)
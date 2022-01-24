package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"ergou/slash"

	"github.com/bwmarrin/discordgo"
)

// launch parameters
var (
	GuildID = flag.String(
		"guildId",	// param name
		"",					// default param value
		"Test guild (server) ID. If left empty, the bot registers commands globally",	// param description
	)
	Token = flag.String("token", "", "Discord bot access token")
	RmCmd = flag.Bool("rmCmd", true, "Remove all commands after shutting down - defaults to true")
)

var discordSession *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	discordSession, err = discordgo.New("Bot " + *Token)
	if err != nil {
		// TODO: handle your errors!
	}
}

func init() {
	discordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := slash.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	discordSession.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is active!")
	})

	err := discordSession.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	// for _, v := range slash.Commands {
	// 	_, err := discordSession.ApplicationCommandCreate(discordSession.State.User.ID, *GuildID, v)
	// 	if err != nil {
	// 		log.Panicf("Failed to create '%v' command: %v", v.Name, err)
	// 	}
	// }

	createdCommands, err := discordSession.ApplicationCommandBulkOverwrite(discordSession.State.User.ID, *GuildID, slash.Commands)
	if err != nil {
		log.Fatalf("Failed to register commands: %v", err)
	}

	defer discordSession.Close()
	
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<- stop
	log.Println("Gracefully shutting down :)")

	if *RmCmd {
		for _, cmd := range createdCommands {
			err := discordSession.ApplicationCommandDelete(discordSession.State.User.ID, *GuildID, cmd.ID)
			if err != nil {
				log.Fatalf("Failed to delete %q command: %v", cmd.Name, err)
			}
		}
	}
}
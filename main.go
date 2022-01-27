package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	GuildID = flag.String(
		"guildId", // param name
		"", // default value
		"Test guld (server) ID", // param description
	)
	Token = flag.String(
		"token",
		"",
		"Discord bot token",
	)
)

func init() { flag.Parse() }

// initialize Discord bot
func init() {
	var err error
	discordSession, err = discordgo.New("Bot " + *Token)
	if err != nil { log.Printf("Failed to initialize Discord bot: %v", err) }
}

// slash command handlers
func init() {
	discordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := CommandHandlers[i.ApplicationCommandData().Name]; ok { h(s, i) }
	})
}

func main() {
	discordSession.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is online & active!")
	})

	err := discordSession.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	createdCommands, err := discordSession.ApplicationCommandBulkOverwrite(discordSession.State.User.ID, *GuildID, Commands)
	if err != nil {
		log.Fatalf("Failed to register commands: %v", err)
	}

	defer discordSession.Close()
	
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<- stop
	log.Println("Gracefully shutting down :)")

	// remove all commands upon shutdown
	log.Println("Removing slash commands...")
	for _, cmd := range createdCommands {
		err := discordSession.ApplicationCommandDelete(discordSession.State.User.ID, *GuildID, cmd.ID)
		if err != nil {
			log.Fatalf("Failed to delete %q command: %v", cmd.Name, err)
		}
	}
}
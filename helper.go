package main

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

func FindUserVoiceChannelFromGuild(
	session *discordgo.Session,
	guildID string,
	userID string,
) (*discordgo.VoiceState, error) {
	guild, err := session.State.Guild(guildID)
	if err != nil {
		return nil, errors.New("error finding guild")
	}

	for _, vs := range guild.VoiceStates {
		if vs.UserID == userID { return vs, nil }
	}

	return nil, errors.New("Unable to find user")
}
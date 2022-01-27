package helpers

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

func FindUserVoiceState(
	session *discordgo.Session,
	guildID string,
	userID string,
) (*discordgo.VoiceState, error) {
	guild, err := session.State.Guild(guildID)
	if err != nil {
		return nil, errors.New("encountered an error while searching for server")
	}

	for _, vs := range guild.VoiceStates {
		if vs.UserID == userID { return vs, nil }
	}

	return nil, errors.New("was unable to find you in an active voice channel")
}
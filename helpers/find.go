package helpers

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

type VoiceState struct { 
	VS *discordgo.VoiceState
	Name string 
}

func FindUserVoiceState(
	session *discordgo.Session,
	guildID string,
	userID string,
) (*VoiceState, error) {
	guild, err := session.State.Guild(guildID)
	if err != nil {
		return nil, errors.New("encountered an error while searching for server")
	}

	for _, vs := range guild.VoiceStates {
		if vs.UserID == userID {
			for _, cn := range guild.Channels {
				if cn.ID == vs.ChannelID { return &VoiceState{vs, cn.Name}, nil }
			}
		}
	}

	return nil, errors.New("was unable to find you in an active voice channel")
}
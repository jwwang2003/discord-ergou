package structs

type SongPkg struct {
	v    *VoiceInstance
	song *Song
}

type Song struct {
	ChannelID string
	User      string
	UserID    string

	Title    string
	Duration string
	VideoID  string
	VideoURL string

	Playlist string
}
package session

type Session struct {
	Identifier           string `json:"id"`
	DeviceName           string `json:"device_name"`
	DeviceIdentifier     string `json:"device_id"`
	Client               string `json:"client"`
	UserName             string `json:"user_name,omitempty"`
	NowPlayingName       string `json:"now_playing_name,omitempty"`
	NowPlayingType       string `json:"now_playing_type,omitempty"`
	NowPlayingIdentifier string `json:"now_playing_id,omitempty"`
	PlayState            string `json:"play_state"`
	Position             string `json:"position,omitempty"`
	SupportsRemote       bool   `json:"supports_remote"`
}

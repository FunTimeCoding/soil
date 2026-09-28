package response

type Session struct {
	Identifier            string      `json:"Id"`
	DeviceName            string      `json:"DeviceName"`
	DeviceIdentifier      string      `json:"DeviceId"`
	Client                string      `json:"Client"`
	UserName              string      `json:"UserName"`
	NowPlayingItem        *NowPlaying `json:"NowPlayingItem"`
	PlayState             *PlayState  `json:"PlayState"`
	SupportsMediaControl  bool        `json:"SupportsMediaControl"`
	SupportsRemoteControl bool        `json:"SupportsRemoteControl"`
}

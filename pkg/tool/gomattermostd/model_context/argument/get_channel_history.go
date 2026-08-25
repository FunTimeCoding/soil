package argument

type GetChannelHistory struct {
	ChannelIdentifier string `json:"channel_id"`
	ChannelName       string `json:"channel_name"`
	Limit             int    `json:"limit"`
	Since             string `json:"since"`
}

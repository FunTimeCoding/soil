package argument

type GetChannel struct {
	ChannelIdentifier string `json:"channel_id"`
	ChannelName       string `json:"channel_name"`
}

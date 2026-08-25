package argument

type PostMessage struct {
	ChannelIdentifier string `json:"channel_id"`
	ChannelName       string `json:"channel_name"`
	Message           string `json:"message"`
}

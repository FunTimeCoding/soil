package argument

type ReplyToThread struct {
	ChannelIdentifier string `json:"channel_id"`
	PostIdentifier    string `json:"post_id"`
	Message           string `json:"message"`
	EmojiName         string `json:"emoji_name"`
}

package argument

type ReplyToThread struct {
	ChannelID string `json:"channel_id"`
	PostID    string `json:"post_id"`
	Message   string `json:"message"`
	EmojiName string `json:"emoji_name"`
}

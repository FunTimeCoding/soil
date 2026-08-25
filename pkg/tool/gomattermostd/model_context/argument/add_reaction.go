package argument

type AddReaction struct {
	PostIdentifier string `json:"post_id"`
	EmojiName      string `json:"emoji_name"`
}

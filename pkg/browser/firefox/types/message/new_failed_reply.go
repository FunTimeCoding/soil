package message

func NewFailedReply(text string) *Reply {
	return &Reply{Error: text}
}

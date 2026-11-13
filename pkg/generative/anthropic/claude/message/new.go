package message

func New(
	role string,
	text string,
) *Message {
	return &Message{Role: role, Text: text}
}

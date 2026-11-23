package message

func New(
	identifier string,
	author string,
	text string,
) *Message {
	return &Message{Identifier: identifier, Author: author, Text: text}
}

package message

func New(
	fromName string,
	toName string,
	body string,
) *Message {
	return &Message{FromName: fromName, ToName: toName, Body: body}
}

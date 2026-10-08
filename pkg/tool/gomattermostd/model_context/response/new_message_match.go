package response

func NewMessageMatch(
	identifier string,
	message string,
	createAt string,
) *MessageMatch {
	return &MessageMatch{
		Identifier: identifier,
		Message:    message,
		CreateAt:   createAt,
	}
}

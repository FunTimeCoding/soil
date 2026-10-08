package response

func NewReply(
	identifier string,
	message string,
	createAt string,
) *Reply {
	return &Reply{
		Identifier: identifier,
		Message:    message,
		CreateAt:   createAt,
	}
}

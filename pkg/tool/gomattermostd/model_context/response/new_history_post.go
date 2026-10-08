package response

func NewHistoryPost(
	identifier string,
	message string,
	createAt string,
) *HistoryPost {
	return &HistoryPost{
		Identifier: identifier,
		Message:    message,
		CreateAt:   createAt,
	}
}

package queue

func NewBroadcast(
	kind string,
	body string,
) *Entry {
	return &Entry{Kind: kind, Body: body}
}

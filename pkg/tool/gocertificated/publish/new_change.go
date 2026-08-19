package publish

func NewChange(
	path string,
	reason string,
	content string,
) *Change {
	return &Change{Path: path, Reason: reason, Content: content}
}

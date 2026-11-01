package entry

func New(
	path string,
	submit string,
) *Entry {
	return &Entry{path: path, submit: submit}
}

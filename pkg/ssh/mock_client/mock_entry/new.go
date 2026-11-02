package mock_entry

func New(
	name string,
	size int64,
) *Entry {
	return &Entry{name: name, size: size}
}

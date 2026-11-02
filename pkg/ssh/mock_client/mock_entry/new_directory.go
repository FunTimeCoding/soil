package mock_entry

func NewDirectory(name string) *Entry {
	return &Entry{name: name, directory: true}
}

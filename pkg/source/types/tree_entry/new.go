package tree_entry

import "time"

func New(
	size int64,
	modified time.Time,
	directory bool,
	linked bool,
) *Entry {
	return &Entry{
		Size:      size,
		Modified:  modified,
		Directory: directory,
		Linked:    linked,
	}
}

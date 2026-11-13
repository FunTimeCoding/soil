package named_entry

import "github.com/funtimecoding/soil/pkg/source/types/tree_entry"

func New(
	name string,
	entry *tree_entry.Entry,
) *Entry {
	return &Entry{name: name, entry: entry}
}

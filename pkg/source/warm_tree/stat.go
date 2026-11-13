package warm_tree

import (
	"github.com/funtimecoding/soil/pkg/source/types/tree_entry"
	"path"
)

func (t *Tree) Stat(p string) (*tree_entry.Entry, bool) {
	e, found := t.entries[path.Clean(p)]

	return e, found
}

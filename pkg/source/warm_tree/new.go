package warm_tree

import (
	"github.com/funtimecoding/soil/pkg/source/types/tree_entry"
	"path/filepath"
)

func New(root string) *Tree {
	result := &Tree{
		root:     filepath.Clean(root),
		entries:  make(map[string]*tree_entry.Entry),
		children: make(map[string][]string),
		content:  make(map[string][]byte),
	}
	result.Refresh()

	return result
}

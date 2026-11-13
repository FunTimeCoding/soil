package warm_tree

import "github.com/funtimecoding/soil/pkg/source/types/tree_entry"

type Tree struct {
	root     string
	entries  map[string]*tree_entry.Entry
	children map[string][]string
	content  map[string][]byte
}

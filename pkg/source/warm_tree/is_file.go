package warm_tree

import "path"

func (t *Tree) IsFile(p string) bool {
	e, found := t.entries[path.Clean(p)]

	return found && !e.Directory
}

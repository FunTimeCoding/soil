package warm_tree

import "path"

func (t *Tree) Children(directory string) []string {
	return t.children[path.Clean(directory)]
}

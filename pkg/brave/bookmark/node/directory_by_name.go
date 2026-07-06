package node

import "github.com/funtimecoding/soil/pkg/brave/bookmark"

func DirectoryByName(
	n *Node,
	name string,
) *Node {
	var result *Node
	WalkUntil(
		n,
		func(o *Node) bool {
			if o.Type == bookmark.DirectoryType && o.Name == name {
				result = o

				return true
			}

			return false
		},
	)

	return result
}

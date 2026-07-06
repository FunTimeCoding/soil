package node

import "github.com/funtimecoding/soil/pkg/errors"

func MustDirectoryByName(
	n *Node,
	name string,
) *Node {
	result := DirectoryByName(n, name)

	if result == nil {
		errors.NotFound(name)

		return Stub()
	}

	return result
}

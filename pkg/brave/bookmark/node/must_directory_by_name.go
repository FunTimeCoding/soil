package node

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func MustDirectoryByName(
	n *Node,
	name string,
) *Node {
	result := DirectoryByName(n, name)

	if result == nil {
		panic(not_found.New("directory", name))
	}

	return result
}

package node

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func MustDirectory(
	n *Node,
	name string,
) *Node {
	result := Directory(n, name)

	if result == nil {
		panic(not_found.New("directory", name))
	}

	return result
}

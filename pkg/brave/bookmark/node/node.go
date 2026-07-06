package node

import "github.com/funtimecoding/soil/pkg/brave/bookmark/file"

type Node struct {
	Type     string
	Name     string
	Link     string
	Children []*Node
	Parent   *Node
	Raw      *file.Node
}

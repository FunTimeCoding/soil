package directory_group

import "github.com/funtimecoding/soil/pkg/brave/bookmark/node"

type Group struct {
	Directory *node.Node
	Links     []*node.Node
}

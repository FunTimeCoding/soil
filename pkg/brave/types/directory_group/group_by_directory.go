package directory_group

import (
	"github.com/funtimecoding/soil/pkg/brave/bookmark/node"
	"github.com/funtimecoding/soil/pkg/brave/constant"
)

func GroupByDirectory(n *node.Node) []*Group {
	var result []*Group
	var traverse func(n *node.Node)
	traverse = func(n *node.Node) {
		if n.Type == constant.DirectoryType {
			var links []*node.Node

			for _, c := range n.Children {
				if c.Type == constant.LinkType {
					links = append(links, c)
				}
			}

			if len(links) > 0 {
				result = append(result, &Group{Directory: n, Links: links})
			}

			for _, c := range n.Children {
				if c.Type == constant.DirectoryType {
					traverse(c)
				}
			}
		}
	}
	traverse(n)

	return result
}

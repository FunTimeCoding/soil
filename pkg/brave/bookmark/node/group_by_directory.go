package node

import "github.com/funtimecoding/soil/pkg/brave/constant"

func GroupByDirectory(n *Node) []*DirectoryGroup {
	var result []*DirectoryGroup
	var traverse func(n *Node)
	traverse = func(n *Node) {
		if n.Type == constant.DirectoryType {
			var links []*Node

			for _, c := range n.Children {
				if c.Type == constant.LinkType {
					links = append(links, c)
				}
			}

			if len(links) > 0 {
				result = append(
					result,
					&DirectoryGroup{Directory: n, Links: links},
				)
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

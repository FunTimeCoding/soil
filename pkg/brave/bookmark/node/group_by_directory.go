package node

import "github.com/funtimecoding/soil/pkg/brave/bookmark"

func GroupByDirectory(n *Node) []*DirectoryGroup {
	var result []*DirectoryGroup
	var traverse func(n *Node)
	traverse = func(n *Node) {
		if n.Type == bookmark.DirectoryType {
			var links []*Node

			for _, c := range n.Children {
				if c.Type == bookmark.LinkType {
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
				if c.Type == bookmark.DirectoryType {
					traverse(c)
				}
			}
		}
	}
	traverse(n)

	return result
}

package node

import "github.com/funtimecoding/soil/pkg/brave/bookmark"

func OnlyLinks(v []*Node) []*Node {
	var result []*Node

	for _, n := range v {
		if n.Type == bookmark.LinkType {
			result = append(result, n)
		}
	}

	return result
}

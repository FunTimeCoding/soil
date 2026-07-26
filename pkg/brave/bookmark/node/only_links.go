package node

import "github.com/funtimecoding/soil/pkg/brave/constant"

func OnlyLinks(v []*Node) []*Node {
	var result []*Node

	for _, n := range v {
		if n.Type == constant.LinkType {
			result = append(result, n)
		}
	}

	return result
}

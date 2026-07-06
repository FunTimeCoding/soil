package node

import "github.com/funtimecoding/soil/pkg/brave/bookmark/file"

func NewSlice(v []*file.Node) []*Node {
	var result []*Node

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}

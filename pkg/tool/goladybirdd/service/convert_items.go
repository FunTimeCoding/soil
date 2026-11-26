package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/node"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/page_tree"
)

func convertItems(
	items []*page_tree.Item,
	shared []string,
	references map[string]string,
) []*node.Node {
	var result []*node.Node

	for _, i := range items {
		uid := ""

		if i.Index != nil && *i.Index < len(shared) {
			uid = fmt.Sprintf("e%d", *i.Index+1)
			references[uid] = shared[*i.Index]
		}

		n := node.New(uid, i.Role, i.Name, i.Value)
		n.Children = convertItems(i.Children, shared, references)
		result = append(result, n)
	}

	return result
}

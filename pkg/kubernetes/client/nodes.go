package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/list"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/node"
)

func (c *Client) Nodes() []*node.Node {
	var result []*node.Node

	for _, l := range c.selectClients(nil) {
		result = append(
			result,
			node.NewSlice(list.Node(l.client, l.context), l.cluster)...,
		)
	}

	return result
}

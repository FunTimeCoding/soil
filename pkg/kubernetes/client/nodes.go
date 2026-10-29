package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/list"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/node"
)

func (c *Client) Nodes() ([]*node.Node, error) {
	var result []*node.Node

	for _, l := range c.selectClients(nil) {
		v, e := list.Node(l.client, l.context)

		if e != nil {
			return nil, e
		}

		result = append(result, node.NewSlice(v, l.cluster)...)
	}

	return result, nil
}

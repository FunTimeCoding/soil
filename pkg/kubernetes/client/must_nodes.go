package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/node"
)

func (c *Client) MustNodes() []*node.Node {
	result, e := c.Nodes()
	errors.PanicOnError(e)

	return result
}

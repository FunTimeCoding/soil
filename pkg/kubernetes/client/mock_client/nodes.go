package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/node"

func (c *Client) Nodes() ([]*node.Node, error) {
	return nil, c.Failure
}

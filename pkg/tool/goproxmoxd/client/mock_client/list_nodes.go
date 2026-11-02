package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListNodes(
	_ context.Context,
	_ *client.ListNodesParams,
) ([]client.Node, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.NodeList, nil
}

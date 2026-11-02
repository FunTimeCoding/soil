package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) GetNodeStatus(
	_ context.Context,
	node string,
	_ *client.GetNodeStatusParams,
) (*client.NodeStatus, error) {
	if e := c.node(node); e != nil {
		return nil, e
	}

	return &client.NodeStatus{}, nil
}

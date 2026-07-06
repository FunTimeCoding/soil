package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
)

func (c *Client) NodeStatus(name string) (*node_status.Status, error) {
	s, okay := c.nodeStatuses[name]

	if !okay {
		return nil, fmt.Errorf("node %q not found", name)
	}

	return s, nil
}

package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
)

func (c *Client) NodeStatus(name string) (*node_status.Status, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	s, okay := c.nodeStatuses[name]

	if !okay {
		return nil, not_found.New(constant.NodeKind, name)
	}

	return s, nil
}

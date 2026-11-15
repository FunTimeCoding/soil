package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Node(name string) (*proxmox.Node, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	n, okay := c.nodes[name]

	if !okay {
		return nil, not_found.New(constant.NodeKind, name)
	}

	return n, nil
}

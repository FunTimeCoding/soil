package mock_client

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Node(name string) (*proxmox.Node, error) {
	n, okay := c.nodes[name]

	if !okay {
		return nil, fmt.Errorf("node %q not found", name)
	}

	return n, nil
}

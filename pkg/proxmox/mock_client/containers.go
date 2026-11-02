package mock_client

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Containers(n *proxmox.Node) (proxmox.Containers, error) {
	nodeContainers, okay := c.containers[n.Name]

	if !okay {
		return nil, fmt.Errorf("node %q not found", n.Name)
	}

	var result proxmox.Containers

	for _, ct := range nodeContainers {
		result = append(result, ct)
	}

	return result, nil
}

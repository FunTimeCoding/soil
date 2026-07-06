package proxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
)

func (c *Client) NodeStatus(name string) (*node_status.Status, error) {
	if name == "" {
		return nil, fmt.Errorf("node name is required")
	}

	n := node_status.New()
	e := Get(c, fmt.Sprintf("/nodes/%s/status", name), &n.Payload)

	if e != nil {
		return nil, e
	}

	return n.Payload, nil
}

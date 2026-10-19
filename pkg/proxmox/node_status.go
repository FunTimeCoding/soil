package proxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
)

func (c *Client) NodeStatus(name string) (*node_status.Status, error) {
	if name == "" {
		return nil, validation.New("node name is required")
	}

	n := node_status.New()
	e := Get(c, fmt.Sprintf("/nodes/%s/status", name), &n.Payload)

	if e != nil {
		return nil, notFound(e, constant.NodeKind, name)
	}

	return n.Payload, nil
}

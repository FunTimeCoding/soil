package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Storage(
	n *proxmox.Node,
	name string,
) (*proxmox.Storage, error) {
	for _, s := range c.storages[n.Name] {
		if s.Name == name {
			return s, nil
		}
	}

	return nil, not_found.New(constant.StorageKind, name)
}

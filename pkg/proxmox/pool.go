package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Pool(name string) (*proxmox.Pool, error) {
	if name == "" {
		return nil, validation.New("pool name is required")
	}

	result, e := c.client.Pool(c.context, name)

	if e != nil {
		return nil, notFound(e, constant.PoolKind, name)
	}

	return result, nil
}

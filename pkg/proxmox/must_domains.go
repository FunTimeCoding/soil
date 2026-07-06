package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustDomains() proxmox.Domains {
	result, e := c.Domains()
	errors.PanicOnError(e)

	return result
}

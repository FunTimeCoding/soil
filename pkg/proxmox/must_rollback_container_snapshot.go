package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustRollbackContainerSnapshot(
	v *proxmox.Container,
	name string,
) *proxmox.Task {
	result, e := c.RollbackContainerSnapshot(v, name)
	errors.PanicOnError(e)

	return result
}

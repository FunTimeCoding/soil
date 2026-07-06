package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustDeleteContainerSnapshot(
	v *proxmox.Container,
	name string,
) *proxmox.Task {
	result, e := c.DeleteContainerSnapshot(v, name)
	errors.PanicOnError(e)

	return result
}

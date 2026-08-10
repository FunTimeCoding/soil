package proxmox

import (
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) ShutdownContainer(v *proxmox.Container) (*proxmox.Task, error) {
	return v.Shutdown(
		c.context,
		constant.ContainerShutdownForce,
		constant.ContainerShutdownTimeout,
	)
}

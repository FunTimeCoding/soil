package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/container_detail"
	"github.com/luthermonson/go-proxmox"
)

func containerDetail(c *proxmox.Container) *container_detail.Detail {
	result := container_detail.New()
	result.Identifier = uint64(c.VMID)
	result.Name = c.Name
	result.Node = c.Node
	result.Status = c.Status
	result.CPUs = c.CPUs
	result.MaxMem = c.MaxMem
	result.MaxDisk = c.MaxDisk
	result.MaxSwap = c.MaxSwap
	result.Uptime = c.Uptime
	result.Tags = c.Tags

	if c.ContainerConfig != nil {
		result.Hostname = c.ContainerConfig.Hostname
		result.Description = c.ContainerConfig.Description
		result.Cores = c.ContainerConfig.Cores

		if c.ContainerConfig.Arch != nil {
			result.Arch = *c.ContainerConfig.Arch
		}

		if c.ContainerConfig.Memory != nil {
			result.Memory = *c.ContainerConfig.Memory
		}
	}

	return result
}

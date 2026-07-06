package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func ContainerDetail(c *proxmox.Container) *server.ContainerDetail {
	result := &server.ContainerDetail{
		Identifier: int64(c.VMID),
		Name:       c.Name,
		Node:       &c.Node,
		Status:     &c.Status,
		Cpus:       new(c.CPUs),
		MaxMem:     new(int64(c.MaxMem)),
		MaxDisk:    new(int64(c.MaxDisk)),
		MaxSwap:    new(int64(c.MaxSwap)),
		Uptime:     new(int64(c.Uptime)),
	}

	if c.Tags != "" {
		result.Tags = &c.Tags
	}

	if c.ContainerConfig != nil {
		if c.ContainerConfig.Hostname != "" {
			result.Hostname = &c.ContainerConfig.Hostname
		}

		if c.ContainerConfig.Description != "" {
			result.Description = &c.ContainerConfig.Description
		}

		if c.ContainerConfig.Arch != nil {
			result.Arch = c.ContainerConfig.Arch
		}

		if c.ContainerConfig.Cores > 0 {
			result.Cores = &c.ContainerConfig.Cores
		}

		if c.ContainerConfig.Memory != nil {
			result.Memory = c.ContainerConfig.Memory
		}
	}

	return result
}

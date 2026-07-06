package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func Container(c proxmox.Container) *server.Container {
	result := &server.Container{
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

	return result
}

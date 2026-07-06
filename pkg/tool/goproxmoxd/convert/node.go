package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func Node(n proxmox.NodeStatus) *server.Node {
	return &server.Node{
		Name:    n.Node,
		Status:  n.Status,
		Cpu:     new(n.CPU),
		MaxCpu:  new(n.MaxCPU),
		Mem:     new(int64(n.Mem)),
		MaxMem:  new(int64(n.MaxMem)),
		Disk:    new(int64(n.Disk)),
		MaxDisk: new(int64(n.MaxDisk)),
		Uptime:  new(int64(n.Uptime)),
	}
}

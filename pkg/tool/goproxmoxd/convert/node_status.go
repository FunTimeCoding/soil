package convert

import (
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func NodeStatus(s *node_status.Status) *server.NodeStatus {
	return &server.NodeStatus{
		Cpu:            new(s.Processor),
		Uptime:         new(s.Uptime),
		KernelVersion:  new(s.KernelVersion),
		ProxmoxVersion: new(s.ProxmoxVersion),
		LoadAverage:    &s.LoadAverage,
		Memory: &server.MemoryInfo{
			Free:  new(s.Memory.Free),
			Total: new(s.Memory.Total),
			Used:  new(s.Memory.Used),
		},
		Swap: &server.MemoryInfo{
			Free:  new(s.Swap.Free),
			Total: new(s.Swap.Total),
			Used:  new(int64(s.Swap.Used)),
		},
		RootFilesystem: &server.MemoryInfo{
			Free:  new(s.RootFilesystem.Free),
			Total: new(s.RootFilesystem.Total),
			Used:  new(s.RootFilesystem.Used),
		},
	}
}

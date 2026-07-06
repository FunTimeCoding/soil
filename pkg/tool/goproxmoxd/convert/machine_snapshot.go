package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func MachineSnapshot(s *proxmox.VirtualMachineSnapshot) *server.Snapshot {
	result := &server.Snapshot{
		Name: s.Name,
	}

	if s.Description != "" {
		result.Description = &s.Description
	}

	if s.Snaptime > 0 {
		result.Timestamp = &s.Snaptime
	}

	if s.Parent != "" {
		result.Parent = &s.Parent
	}

	return result
}

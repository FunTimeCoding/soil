package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func ContainerSnapshot(s *proxmox.ContainerSnapshot) *server.Snapshot {
	result := &server.Snapshot{
		Name: s.Name,
	}

	if s.Description != "" {
		result.Description = &s.Description
	}

	if s.SnapshotCreationTime > 0 {
		result.Timestamp = &s.SnapshotCreationTime
	}

	if s.Parent != "" {
		result.Parent = &s.Parent
	}

	return result
}

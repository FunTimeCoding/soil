package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func ContainerSnapshots(items []*proxmox.ContainerSnapshot) []*server.Snapshot {
	result := make([]*server.Snapshot, len(items))

	for i, s := range items {
		result[i] = ContainerSnapshot(s)
	}

	return result
}

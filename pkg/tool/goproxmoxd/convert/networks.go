package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func Networks(items proxmox.NodeNetworks) []*server.Network {
	result := make([]*server.Network, len(items))

	for i, n := range items {
		result[i] = Network(*n)
	}

	return result
}

package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func Machines(items proxmox.VirtualMachines) []*server.Machine {
	result := make([]*server.Machine, len(items))

	for i, v := range items {
		result[i] = Machine(*v)
	}

	return result
}

package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func Containers(items proxmox.Containers) []*server.Container {
	result := make([]*server.Container, len(items))

	for i, c := range items {
		result[i] = Container(*c)
	}

	return result
}

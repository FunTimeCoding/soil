package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Disk(d *virtual_disk.Disk) *server.VirtualDisk {
	return &server.VirtualDisk{
		Identifier: d.Identifier,
		Name:       d.Name,
		Size:       d.Size,
	}
}

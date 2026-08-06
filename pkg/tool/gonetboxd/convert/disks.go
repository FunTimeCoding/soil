package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Disks(v []*virtual_disk.Disk) []*server.VirtualDisk {
	result := make([]*server.VirtualDisk, 0, len(v))

	for _, d := range v {
		result = append(result, Disk(d))
	}

	return result
}

package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
)

func New() *Client {
	result := &Client{
		interfaces: make(map[int32][]*network.Interface),
		disks:      make(map[int32][]*virtual_disk.Disk),
		addresses:  make(map[string][]*internet_address.Address),
	}
	result.failures = failure.New(result.Racks)

	return result
}

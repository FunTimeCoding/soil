package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		interfaces: make(map[int32][]*network.Interface),
		outlets:    make(map[int32][]*power_outlet.Outlet),
		ports:      make(map[int32][]*power_port.Port),
		disks:      make(map[int32][]*virtual_disk.Disk),
		addresses:  make(map[string][]*internet_address.Address),
	}
	result.failures = failure.New(result.Racks)
	t.Cleanup(result.Close)

	return result
}

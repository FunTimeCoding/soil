package mock_inventory_source

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"

type Client struct {
	devices  []*client.Device
	machines []*client.VirtualMachine
	physical []client.PhysicalAddressOwner
}

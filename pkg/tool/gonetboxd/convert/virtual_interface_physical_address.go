package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/netbox-community/go-netbox/v4"
)

func VirtualInterfacePhysicalAddress(
	v *netbox.VMInterface,
	address string,
) *server.PhysicalAddress {
	name := v.GetName()

	return &server.PhysicalAddress{
		Identifier: v.GetId(),
		Address:    address,
		Interface:  &name,
	}
}

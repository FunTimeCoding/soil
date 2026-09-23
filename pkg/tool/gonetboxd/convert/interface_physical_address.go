package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func InterfacePhysicalAddress(v *network.Interface) *server.PhysicalAddress {
	name := v.Name

	return &server.PhysicalAddress{
		Identifier: v.Identifier,
		Address:    v.PhysicalAddress.String(),
		Interface:  &name,
	}
}

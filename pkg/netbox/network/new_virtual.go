package network

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/network"
	networkConstant "github.com/funtimecoding/soil/pkg/network/constant"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func NewVirtual(i *netbox.VMInterface) *Interface {
	var name string

	if s := i.GetName(); s != "" {
		name = i.GetName()
	} else {
		name = constant.NoName
	}

	var h net.HardwareAddr

	if s := i.GetMacAddress(); s != "" {
		h = network.PhysicalAddress(s)
	} else {
		h = networkConstant.NullPhysicalAddress
	}

	return &Interface{
		Identifier:      i.GetId(),
		Name:            name,
		Description:     i.GetDescription(),
		Type:            constant.InterfaceVirtual,
		PhysicalAddress: h,
	}
}

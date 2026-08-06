package network

import (
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/network/constant"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func NewVirtual(i *netbox.VMInterface) *Interface {
	var name string

	if s := i.GetName(); s != "" {
		name = i.GetName()
	} else {
		name = netboxConstant.NoName
	}

	var h net.HardwareAddr

	if s := i.GetMacAddress(); s != "" {
		h = network.PhysicalAddress(s)
	} else {
		h = constant.NullPhysicalAddress
	}

	return &Interface{
		Identifier:      i.GetId(),
		Name:            name,
		Description:     i.GetDescription(),
		Type:            netboxConstant.InterfaceVirtual,
		PhysicalAddress: h,
	}
}

package network

import (
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func New(i *netbox.Interface) *Interface {
	var name string

	if s := i.GetName(); s != "" {
		name = i.GetName()
	} else {
		name = NoName
	}

	var h net.HardwareAddr

	if s := i.GetMacAddress(); s != "" {
		h = network.PhysicalAddress(s)
	} else {
		h = network.NullPhysicalAddress
	}

	var t netbox.InterfaceTypeValue

	if i.Type.Value != nil {
		t = *i.Type.Value
		validateType(t)
	}

	return &Interface{
		Identifier:       i.GetId(),
		Name:             name,
		Description:      i.GetDescription(),
		Type:             t,
		PhysicalAddress:  h,
		Contexts:         i.GetVdcs(),
		VirtualNetworks:  i.GetTaggedVlans(),
		WirelessNetworks: i.GetWirelessLans(),
		Raw:              i,
	}
}

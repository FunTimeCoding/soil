package network

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/network"
)

func (i *Interface) Format(f *option.Format) string {
	s := status.New(f).String(i.formatName(f), i.formatType()).RawList(i.Raw)

	if i.PhysicalAddress.String() != network.NullPhysicalAddress.String() {
		s.String(i.PhysicalAddress.String())
	}

	s.String(i.Description)

	for _, c := range i.Contexts {
		s.Line("  Context: %+v", c)
	}

	for _, n := range i.VirtualNetworks {
		s.Line("  Virtual network: %+v", n)
	}

	for _, n := range i.WirelessNetworks {
		s.Line("  Wireless network: %+v", n)
	}

	return s.Format()
}

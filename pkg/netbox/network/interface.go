package network

import (
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

type Interface struct {
	Identifier       int32
	Name             string
	Description      string
	PhysicalAddress  net.HardwareAddr
	Type             netbox.InterfaceTypeValue
	EthernetPowerMode string
	EthernetPowerType string
	Contexts         []netbox.VirtualDeviceContext
	VirtualNetworks  []netbox.VLAN
	WirelessNetworks []netbox.WirelessLAN
	Raw              *netbox.Interface
}

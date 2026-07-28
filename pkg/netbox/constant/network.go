package constant

import "github.com/netbox-community/go-netbox/v4"

const (
	Eth0 = "eth0"
	Eth1 = "eth1"

	InterfaceVirtual   = "virtual"
	Interface1000BaseT = "1000base-t"
)

var InterfaceTypes = []netbox.InterfaceTypeValue{
	InterfaceVirtual,
	Interface1000BaseT,
}

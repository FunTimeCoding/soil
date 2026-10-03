package unit

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
)

func newVirtualInterface() *network.Interface {
	return network.New(
		&netbox.Interface{
			Name: constant.Eth0,
			Type: netbox.InterfaceType{
				Value: new(netbox.InterfaceTypeValue(constant.InterfaceVirtual)),
			},
		},
	)
}

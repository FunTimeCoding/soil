package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/network_interface"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func NetworkInterface(v *network_interface.Interface) *server.NetworkInterface {
	return &server.NetworkInterface{
		Device:     v.Device,
		Status:     v.Status,
		Media:      v.Media,
		MacAddress: v.MacAddress,
		Mtu:        v.Mtu,
		Addresses:  v.Addresses,
	}
}

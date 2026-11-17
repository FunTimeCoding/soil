package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func PowerOutlet(o *power_outlet.Outlet) *server.PowerOutlet {
	result := &server.PowerOutlet{
		Identifier: o.Identifier,
		Name:       o.Name,
		Device:     o.Device,
	}

	if o.ConnectedDevice != "" {
		result.ConnectedDevice = &o.ConnectedDevice
	}

	if o.ConnectedPort != "" {
		result.ConnectedPort = &o.ConnectedPort
	}

	return result
}

package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Devices(v []*device.Device) []*server.Device {
	result := make([]*server.Device, 0, len(v))

	for _, d := range v {
		result = append(result, Device(d))
	}

	return result
}

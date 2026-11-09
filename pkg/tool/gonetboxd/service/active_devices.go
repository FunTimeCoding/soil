package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func activeDevices(v []*device.Device) []*device.Device {
	var result []*device.Device

	for _, d := range v {
		if d.Status == constant.StatusActive {
			result = append(result, d)
		}
	}

	return result
}

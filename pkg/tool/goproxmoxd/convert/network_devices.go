package convert

import (
	"github.com/funtimecoding/soil/pkg/proxmox/network_device"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func NetworkDevices(items []*network_device.Device) []server.NetworkDevice {
	result := make([]server.NetworkDevice, len(items))

	for i, v := range items {
		result[i] = *NetworkDevice(v)
	}

	return result
}

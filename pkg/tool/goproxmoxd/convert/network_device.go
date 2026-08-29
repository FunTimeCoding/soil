package convert

import (
	"github.com/funtimecoding/soil/pkg/proxmox/network_device"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func NetworkDevice(v *network_device.Device) *server.NetworkDevice {
	result := &server.NetworkDevice{Name: v.Name}

	if v.Model != "" {
		result.Model = &v.Model
	}

	if v.Interface != "" {
		result.Interface = &v.Interface
	}

	if v.HardwareAddress != "" {
		result.HardwareAddress = &v.HardwareAddress
	}

	if v.Bridge != "" {
		result.Bridge = &v.Bridge
	}

	if v.Vlan != 0 {
		result.Vlan = &v.Vlan
	}

	return result
}

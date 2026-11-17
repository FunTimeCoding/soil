package power_outlet

import "github.com/funtimecoding/soil/pkg/netbox/constant"

func peer(peers []any) (string, string) {
	if len(peers) == 0 {
		return "", ""
	}

	port, _ := peers[0].(map[string]any)
	name, _ := port[constant.NameField].(string)
	device, _ := port[constant.DeviceField].(map[string]any)
	deviceName, _ := device[constant.NameField].(string)

	return deviceName, name
}

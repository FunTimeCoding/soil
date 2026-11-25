package internet_address

import "github.com/funtimecoding/soil/pkg/netbox/constant"

func holder(assigned any) (string, string, string) {
	object, okay := assigned.(map[string]any)

	if !okay {
		return "", "", ""
	}

	return text(object, constant.NameField),
		nestedName(object, constant.DeviceField),
		nestedName(object, constant.VirtualMachineField)
}

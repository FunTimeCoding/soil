package device_role

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.DeviceRole) *Role {
	return &Role{
		Identifier:     v.GetId(),
		Name:           v.GetName(),
		VirtualMachine: v.GetVmRole(),
		Raw:            v,
	}
}

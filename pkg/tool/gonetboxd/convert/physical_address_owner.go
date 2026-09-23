package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func PhysicalAddressOwner(
	v *physical_address.Address,
) *server.PhysicalAddressOwner {
	result := &server.PhysicalAddressOwner{Address: v.Address.String()}

	if v.Interface != nil {
		kind := constant.DeviceAddress
		result.ObjectKind = &kind
		result.ObjectIdentifier = &v.Interface.Device.Id
		result.ObjectName = v.Interface.Device.Name.Get()
		result.Interface = &v.Interface.Name

		return result
	}

	if v.VirtualInterface != nil {
		kind := constant.VirtualMachineAddress
		result.ObjectKind = &kind
		result.ObjectIdentifier = &v.VirtualInterface.VirtualMachine.Id
		result.ObjectName = &v.VirtualInterface.VirtualMachine.Name
		result.Interface = &v.VirtualInterface.Name
	}

	return result
}

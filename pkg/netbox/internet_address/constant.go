package internet_address

import "github.com/funtimecoding/soil/pkg/netbox/constant"

const (
	// DeviceAddress when assigning an IP address to a device
	DeviceAddress = "dcim.device"

	// VirtualMachineAddress when assigning an IP address to a virtual machine
	VirtualMachineAddress = "virtualization.vm"

	// VirtualInterfaceAddress when assigning an IP address to a VM interface
	VirtualInterfaceAddress = "virtualization.vminterface"

	// SubnetAddress  when assigning an IP address to a prefix/subnet
	SubnetAddress = "ipam.prefix"

	NoObjectType = "no object type"

	fixtureAddress = "192.168.0.1/24"
)

var ObjectTypes = []string{
	DeviceAddress,
	constant.InterfaceAddress,
	VirtualMachineAddress,
	VirtualInterfaceAddress,
	SubnetAddress,
}

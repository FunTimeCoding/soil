package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

type InterfaceClient interface {
	CreateInterface(
		d *device.Device,
		name string,
		t netbox.InterfaceTypeValue,
	) (*network.Interface, error)
	CreatePhysicalInterface(
		a net.HardwareAddr,
		description string,
		i *network.Interface,
	) (*physical_address.Address, error)
	UpdateInterface(
		d *device.Device,
		name string,
		t netbox.InterfaceTypeValue,
		h net.HardwareAddr,
	) (*network.Interface, error)
	PhysicalAddresses() ([]*physical_address.Address, error)
	DeviceInterfaceByName(
		d *device.Device,
		name string,
	) (*network.Interface, error)
	DeviceInterfaces(device int32) ([]*network.Interface, error)
	AssignPhysicalToInterface(
		p *physical_address.Address,
		i *network.Interface,
	) (*physical_address.Address, error)
	CreateInterfacePhysical(
		d *device.Device,
		name string,
		t netbox.InterfaceTypeValue,
		h net.HardwareAddr,
	) (*network.Interface, error)
	CreatePhysical(
		a net.HardwareAddr,
		description string,
	) (*physical_address.Address, error)
	DeleteInterface(identifier int32) error
	DeleteInterfaceByName(
		d *device.Device,
		name string,
	) error
	DeletePhysical(identifier int32) error
	DeviceInterfaceExists(
		d *device.Device,
		name string,
	) (bool, error)
	EnsurePhysical(
		h net.HardwareAddr,
	) (*physical_address.Address, error)
	InterfaceAddresses(interfaceIdentifier int32) ([]*internet_address.Address, error)
	MustCreateInterface(
		d *device.Device,
		name string,
		t netbox.InterfaceTypeValue,
	) *network.Interface
	MustCreatePhysical(
		a net.HardwareAddr,
		description string,
	) *physical_address.Address
	MustDeleteInterface(identifier int32)
	MustDeviceInterfaceByName(
		d *device.Device,
		name string,
	) *network.Interface
	MustDeviceInterfaces(device int32) []*network.Interface
	MustInterfaceAddresses(interfaceIdentifier int32) []*internet_address.Address
	MustPhysicalAddress(a net.HardwareAddr) *physical_address.Address
	MustPhysicalAddresses() []*physical_address.Address
	MustPhysicalAddressesByHardware(a net.HardwareAddr) []*physical_address.Address
	MustUpdateInterface(
		d *device.Device,
		name string,
		t netbox.InterfaceTypeValue,
		h net.HardwareAddr,
	) *network.Interface
	MustUpdatePhysicalAddress(
		a *physical_address.Address,
		i *network.Interface,
	) *physical_address.Address
	PhysicalAddress(a net.HardwareAddr) (*physical_address.Address, error)
	PhysicalAddressesByHardware(a net.HardwareAddr) ([]*physical_address.Address, error)
	PhysicalExists(a net.HardwareAddr) (bool, error)
	UpdatePhysicalAddress(
		a *physical_address.Address,
		i *network.Interface,
	) (*physical_address.Address, error)
}

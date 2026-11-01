package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

type DeviceClient interface {
	CreateDevice(
		name string,
		o *device_role.Role,
		tags []string,
		y *device_type.Type,
		s *site.Site,
		n *tenant.Tenant,
	) (*device.Device, error)
	DeviceByName(n string) (*device.Device, error)
	Devices() ([]*device.Device, error)
	DevicesByMatch(s string) ([]*device.Device, error)
	RenameDevice(
		name string,
		newName string,
	) (*device.Device, error)
	SetDeviceDescription(
		name string,
		description string,
	) (*device.Device, error)
	SetDeviceLocation(
		name string,
		locationName string,
	) (*device.Device, error)
	SetDevicePlatform(
		name string,
		platformName string,
	) (*device.Device, error)
	SetDevicePrimaryAddress(
		name string,
		address string,
	) (*device.Device, error)
	SetDeviceSerial(
		name string,
		serial string,
	) (*device.Device, error)
	SetDeviceStatus(
		name string,
		status string,
	) (*device.Device, error)
	SetDeviceTenant(
		name string,
		tenantName string,
	) (*device.Device, error)
	DeleteDevice(identifier int32) error
	DeviceByNames(n []string) (*device.Device, error)
	DeviceExists(name string) (bool, error)
	DevicesByName(s string) ([]*device.Device, error)
	MustCreateDevice(
		name string,
		o *device_role.Role,
		tags []string,
		y *device_type.Type,
		s *site.Site,
		n *tenant.Tenant,
	) *device.Device
	MustDeviceByName(n string) *device.Device
	MustDevices() []*device.Device
	MustDevicesByMatch(s string) []*device.Device
	MustDevicesByName(s string) []*device.Device
	UnsetSerial(device string) (*device.Device, error)
	UpdateSerial(
		device string,
		serial string,
	) (*device.Device, error)
}

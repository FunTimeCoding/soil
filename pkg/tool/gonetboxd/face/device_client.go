package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

type DeviceClient interface {
	AnsibleInventoryMarkup(
		x context.Context,
		field string,
	) (string, error)
	CreateCable(
		x context.Context,
		deviceA string,
		interfaceA string,
		deviceB string,
		interfaceB string,
	) (client.Cable, error)
	CreateDevice(
		x context.Context,
		name string,
		role string,
		deviceType string,
		site string,
		tenant *string,
	) (client.Device, error)
	CreateDeviceRole(
		x context.Context,
		name string,
		virtualMachine bool,
	) (client.DeviceRole, error)
	CreateDeviceType(
		x context.Context,
		model string,
		manufacturer string,
	) (client.DeviceType, error)
	CreateInterface(
		x context.Context,
		device string,
		name string,
		interfaceType string,
		physicalAddress string,
	) (client.Interface, error)
	GetDevice(
		x context.Context,
		name string,
	) (client.Device, error)
	ListCables(
		x context.Context,
	) ([]*client.Cable, error)
	ListDeviceRoles(
		x context.Context,
	) ([]*client.DeviceRole, error)
	ListDeviceTypes(
		x context.Context,
	) ([]*client.DeviceType, error)
	ListDevices(
		x context.Context,
		query *string,
	) ([]*client.Device, error)
	ListInterfaces(
		x context.Context,
		name string,
	) ([]*client.Interface, error)
	UpdateDevice(
		x context.Context,
		name string,
		body client.UpdateDeviceRequest,
	) (client.Device, error)
}

package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
)

type PowerClient interface {
	CreatePowerCable(
		o *power_outlet.Outlet,
		p *power_port.Port,
	) (*cable.Cable, error)
	CreatePowerOutlet(
		d *device.Device,
		name string,
	) (*power_outlet.Outlet, error)
	CreatePowerPort(
		d *device.Device,
		name string,
	) (*power_port.Port, error)
	DevicePowerOutlet(
		d *device.Device,
		name string,
	) (*power_outlet.Outlet, error)
	DevicePowerOutlets(device int32) ([]*power_outlet.Outlet, error)
	DevicePowerPort(
		d *device.Device,
		name string,
	) (*power_port.Port, error)
	DevicePowerPorts(device int32) ([]*power_port.Port, error)
}

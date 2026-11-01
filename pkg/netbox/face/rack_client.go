package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/power_feed"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/netbox/power_panel"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
	"github.com/funtimecoding/soil/pkg/netbox/rack_role"
	"github.com/funtimecoding/soil/pkg/netbox/rack_type"
)

type RackClient interface {
	DevicesByRack(i int32) ([]*device.Device, error)
	MustDevicesByRack(i int32) []*device.Device
	MustPowerFeeds() []*power_feed.Feed
	MustPowerOutlets() []*power_outlet.Outlet
	MustPowerPanels() []*power_panel.Panel
	MustPowerPorts() []*power_port.Port
	MustRackByName(n string) *rack.Rack
	MustRackRoles() []*rack_role.Role
	MustRackTypes() []*rack_type.Type
	MustRacks() []*rack.Rack
	PowerFeeds() ([]*power_feed.Feed, error)
	PowerOutlets() ([]*power_outlet.Outlet, error)
	PowerPanels() ([]*power_panel.Panel, error)
	PowerPorts() ([]*power_port.Port, error)
	RackByName(n string) (*rack.Rack, error)
	RackRoles() ([]*rack_role.Role, error)
	RackTypes() ([]*rack_type.Type, error)
	Racks() ([]*rack.Rack, error)
	RacksByName(n string) ([]*rack.Rack, error)
}

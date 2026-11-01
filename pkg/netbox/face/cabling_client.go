package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/netbox/console_port"
	"github.com/funtimecoding/soil/pkg/netbox/console_server_port"
	"github.com/funtimecoding/soil/pkg/netbox/front_port"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/rear_port"
)

type CablingClient interface {
	Cables() ([]*cable.Cable, error)
	CreateCable(
		a *network.Interface,
		b *network.Interface,
	) (*cable.Cable, error)
	ConsolePorts() ([]*console_port.Port, error)
	ConsoleServerPorts() ([]*console_server_port.Port, error)
	DeviceCables(device string) ([]*cable.Cable, error)
	FrontPorts() ([]*front_port.Port, error)
	MustCables() []*cable.Cable
	MustConsolePorts() []*console_port.Port
	MustConsoleServerPorts() []*console_server_port.Port
	MustDeviceCables(device string) []*cable.Cable
	MustFrontPorts() []*front_port.Port
	MustRearPorts() []*rear_port.Port
	RearPorts() ([]*rear_port.Port, error)
}

package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func PowerPort(p *power_port.Port) *server.PowerPort {
	return &server.PowerPort{
		Identifier: p.Identifier,
		Name:       p.Name,
		Device:     p.Device,
	}
}

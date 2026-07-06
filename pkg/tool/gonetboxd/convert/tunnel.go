package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Tunnel(t *tunnel.Tunnel) *server.Tunnel {
	result := &server.Tunnel{Identifier: t.Identifier, Name: t.Name}

	if t.Raw.Encapsulation.Value != nil {
		result.Encapsulation = new(string(*t.Raw.Encapsulation.Value))
	}

	if t.Raw.Group.IsSet() && t.Raw.Group.Get() != nil {
		result.Group = new(t.Raw.Group.Get().GetName())
	}

	return result
}

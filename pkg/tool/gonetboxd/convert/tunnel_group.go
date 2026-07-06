package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func TunnelGroup(g *tunnel_group.Group) *server.TunnelGroup {
	return &server.TunnelGroup{Identifier: g.Identifier, Name: g.Name}
}

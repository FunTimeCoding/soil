package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func TunnelGroups(v []*tunnel_group.Group) []*server.TunnelGroup {
	result := make([]*server.TunnelGroup, 0, len(v))

	for _, g := range v {
		result = append(result, TunnelGroup(g))
	}

	return result
}

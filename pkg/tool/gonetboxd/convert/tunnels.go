package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Tunnels(v []*tunnel.Tunnel) []*server.Tunnel {
	result := make([]*server.Tunnel, 0, len(v))

	for _, t := range v {
		result = append(result, Tunnel(t))
	}

	return result
}

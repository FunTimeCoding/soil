package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func TunnelTerminations(v []*tunnel_termination.Termination) []*server.TunnelTermination {
	result := make([]*server.TunnelTermination, 0, len(v))

	for _, t := range v {
		result = append(result, TunnelTermination(t))
	}

	return result
}

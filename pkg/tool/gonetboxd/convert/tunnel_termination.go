package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func TunnelTermination(t *tunnel_termination.Termination) *server.TunnelTermination {
	result := &server.TunnelTermination{Identifier: t.Identifier}

	if t.Role != "" {
		result.Role = &t.Role
	}

	if t.TerminationType != "" {
		result.TerminationType = &t.TerminationType
	}

	if t.TerminationIdentifier != 0 {
		result.TerminationIdentifier = &t.TerminationIdentifier
	}

	if t.Raw.Tunnel.Name != "" {
		result.Tunnel = &t.Raw.Tunnel.Name
	}

	return result
}

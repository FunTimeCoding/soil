package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/state"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func State(v *state.State) *server.State {
	return &server.State{
		Label:              v.Label,
		Description:        v.Description,
		Interface:          v.Interface,
		Protocol:           v.Protocol,
		Direction:          v.Direction,
		SourceAddress:      v.SourceAddress,
		SourcePort:         v.SourcePort,
		DestinationAddress: v.DestinationAddress,
		DestinationPort:    v.DestinationPort,
		State:              v.State,
		Age:                v.Age,
		Expires:            v.Expires,
		Packets:            v.Packets,
		Bytes:              v.Bytes,
		Rule:               v.Rule,
	}
}

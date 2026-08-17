package state

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func New(v response.State) *State {
	return &State{
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

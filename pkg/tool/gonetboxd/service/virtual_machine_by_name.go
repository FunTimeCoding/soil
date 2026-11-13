package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) virtualMachineByName(
	name string,
) (*virtual_machine.Machine, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*virtual_machine.Machine, error) {
			return s.client.VirtualMachine(name)
		},
	)
}

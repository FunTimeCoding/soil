package service

import (
	"context"
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateVirtualTunnelTermination(
	machine string,
	tunnel string,
	interfaceName string,
	role string,
) (*tunnel_termination.Termination, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*tunnel_termination.Termination, error) {
			t, e := s.client.Tunnel(tunnel)

			if e != nil {
				return nil, e
			}

			m, e := s.client.VirtualMachine(machine)

			if e != nil {
				return nil, e
			}

			i, e := s.client.VirtualMachineInterface(m, interfaceName)

			if e != nil {
				return nil, e
			}

			return s.client.CreateTunnelTermination(
				t,
				netboxConstant.VirtualInterfaceAddress,
				int64(i.GetId()),
				role,
			)
		},
	)
}

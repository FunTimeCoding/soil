package service

import (
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
)

func (s *Service) CreateVirtualTunnelTermination(
	machine string,
	tunnel string,
	interfaceName string,
	role string,
) (*tunnel_termination.Termination, error) {
	return guarded(
		s.guard,
		func() (*tunnel_termination.Termination, error) {
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

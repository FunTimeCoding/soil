package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
)

func (s *Service) CreateVirtualTunnelTermination(
	machine string,
	tunnel string,
	interfaceName string,
	role string,
) (*tunnel_termination.Termination, error) {
	t, e := s.client.TunnelByName(tunnel)

	if e != nil {
		return nil, e
	}

	m, e := s.client.VirtualMachineByName(machine)

	if e != nil {
		return nil, e
	}

	i, e := s.client.VirtualMachineInterfaceByName(m, interfaceName)

	if e != nil {
		return nil, e
	}

	return s.client.CreateTunnelTermination(
		t,
		constant.VirtualInterfaceAddress,
		int64(i.GetId()),
		role,
	)
}

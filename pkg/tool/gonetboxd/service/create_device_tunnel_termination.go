package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
)

func (s *Service) CreateDeviceTunnelTermination(
	device string,
	tunnel string,
	interfaceName string,
	role string,
) (*tunnel_termination.Termination, error) {
	t, e := s.client.TunnelByName(tunnel)

	if e != nil {
		return nil, e
	}

	d, e := s.client.DeviceByName(device)

	if e != nil {
		return nil, e
	}

	i, e := s.client.DeviceInterfaceByName(d, interfaceName)

	if e != nil {
		return nil, e
	}

	return s.client.CreateTunnelTermination(
		t,
		constant.InterfaceAddress,
		int64(i.Identifier),
		role,
	)
}

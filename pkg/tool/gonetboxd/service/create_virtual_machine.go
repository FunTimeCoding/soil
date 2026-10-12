package service

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (s *Service) CreateVirtualMachine(
	name string,
	cluster string,
) (*virtual_machine.Machine, error) {
	c, e := s.client.ClusterByName(cluster)

	if e != nil {
		return nil, e
	}

	return s.client.CreateVirtualMachine(name, c)
}

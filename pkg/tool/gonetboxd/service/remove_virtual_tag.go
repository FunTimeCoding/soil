package service

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (s *Service) RemoveVirtualTag(
	name string,
	tag string,
) (*virtual_machine.Machine, error) {
	return s.client.RemoveVirtualTag(name, tag)
}

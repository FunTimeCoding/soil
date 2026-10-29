package service

import "github.com/funtimecoding/soil/pkg/opnsense/network_interface"

func (s *Service) Interfaces() ([]*network_interface.Interface, error) {
	return s.source.Interfaces()
}

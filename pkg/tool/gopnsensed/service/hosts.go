package service

import "github.com/funtimecoding/soil/pkg/opnsense/host"

func (s *Service) Hosts(phrase string) ([]*host.Host, error) {
	return s.source.Hosts(phrase)
}

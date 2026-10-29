package service

import "github.com/funtimecoding/soil/pkg/netbox/prefix"

func (s *Service) Prefixes() ([]*prefix.Prefix, error) {
	return s.client.Prefixes()
}

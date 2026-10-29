package service

import "github.com/funtimecoding/soil/pkg/netbox/platform"

func (s *Service) CreatePlatform(name string) (*platform.Platform, error) {
	return s.client.CreatePlatform(name)
}

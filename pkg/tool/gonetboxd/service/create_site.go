package service

import "github.com/funtimecoding/soil/pkg/netbox/site"

func (s *Service) CreateSite(name string) (*site.Site, error) {
	return s.client.CreateSite(name)
}

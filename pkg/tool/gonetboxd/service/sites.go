package service

import "github.com/funtimecoding/soil/pkg/netbox/site"

func (s *Service) Sites() ([]*site.Site, error) {
	return s.client.Sites()
}

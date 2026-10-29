package service

import "github.com/funtimecoding/soil/pkg/netbox/tag"

func (s *Service) CreateTag(name string) (*tag.Tag, error) {
	return s.client.CreateTag(name)
}

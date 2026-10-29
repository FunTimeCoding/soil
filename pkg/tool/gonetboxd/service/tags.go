package service

import "github.com/funtimecoding/soil/pkg/netbox/tag"

func (s *Service) Tags() ([]*tag.Tag, error) {
	return s.client.Tags()
}

package service

import "github.com/funtimecoding/soil/pkg/netbox/configuration_context"

func (s *Service) ConfigurationContextByName(
	n string,
) (*configuration_context.Context, error) {
	return s.client.ConfigurationContextByName(n)
}

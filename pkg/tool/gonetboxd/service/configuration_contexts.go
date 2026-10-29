package service

import "github.com/funtimecoding/soil/pkg/netbox/configuration_context"

func (s *Service) ConfigurationContexts() ([]*configuration_context.Context, error) {
	return s.client.ConfigurationContexts()
}

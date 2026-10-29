package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
)

func (s *Service) CreateConfigurationContext(
	name string,
	weight int32,
	payload map[string]any,
	a *configuration_assignment.Assignment,
) (*configuration_context.Context, error) {
	return s.client.CreateConfigurationContext(name, weight, payload, a)
}

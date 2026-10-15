package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/context_change"
)

func (s *Service) UpdateConfigurationContext(
	name string,
	c *context_change.Change,
) (*configuration_context.Context, error) {
	if _, e := s.client.ConfigurationContextByName(name); e != nil {
		return nil, e
	}

	if c.Payload != nil {
		if _, e := s.client.SetConfigurationContextPayload(
			name,
			c.Payload,
		); e != nil {
			return nil, e
		}
	}

	if c.Weight != nil {
		if _, e := s.client.SetConfigurationContextWeight(
			name,
			*c.Weight,
		); e != nil {
			return nil, e
		}
	}

	if c.Assignment != nil {
		if _, e := s.client.SetConfigurationContextAssignment(
			name,
			c.Assignment,
		); e != nil {
			return nil, e
		}
	}

	return s.client.ConfigurationContextByName(name)
}

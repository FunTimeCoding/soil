package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func ConfigurationContext(
	v *configuration_context.Context,
) *server.ConfigurationContext {
	result := &server.ConfigurationContext{
		Identifier: v.Identifier,
		Name:       v.Name,
		Weight:     v.Weight,
		Active:     v.Active,
		Payload:    v.Payload,
		Assignment: server.ConfigurationAssignment{
			Roles:     &v.Roles,
			Clusters:  &v.Clusters,
			Platforms: &v.Platforms,
			Sites:     &v.Sites,
			Tenants:   &v.Tenants,
			Tags:      &v.Tags,
		},
	}

	if result.Payload == nil {
		result.Payload = map[string]any{}
	}

	if v.Description != "" {
		result.Description = &v.Description
	}

	return result
}

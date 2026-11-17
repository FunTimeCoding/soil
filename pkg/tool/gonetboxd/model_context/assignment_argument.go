package model_context

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
)

func assignmentArgument(
	a argument.ConfigurationContext,
) *configuration_assignment.Assignment {
	if a.Roles == nil && a.Clusters == nil && a.Platforms == nil &&
		a.Sites == nil && a.Tenants == nil && a.Tags == nil {
		return nil
	}

	result := configuration_assignment.New()
	result.Roles = given(a.Roles)
	result.Clusters = given(a.Clusters)
	result.Platforms = given(a.Platforms)
	result.Sites = given(a.Sites)
	result.Tenants = given(a.Tenants)
	result.Tags = given(a.Tags)

	return result
}

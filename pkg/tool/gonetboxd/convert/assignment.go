package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Assignment(
	v *server.ConfigurationAssignment,
) *configuration_assignment.Assignment {
	result := configuration_assignment.New()

	if v == nil {
		return result
	}

	if v.Roles != nil {
		result.Roles = *v.Roles
	}

	if v.Clusters != nil {
		result.Clusters = *v.Clusters
	}

	if v.Platforms != nil {
		result.Platforms = *v.Platforms
	}

	if v.Sites != nil {
		result.Sites = *v.Sites
	}

	if v.Tenants != nil {
		result.Tenants = *v.Tenants
	}

	if v.Tags != nil {
		result.Tags = *v.Tags
	}

	return result
}

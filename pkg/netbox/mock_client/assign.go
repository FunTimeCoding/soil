package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
)

func assign(
	v *configuration_context.Context,
	a *configuration_assignment.Assignment,
) {
	v.Roles = a.Roles
	v.Clusters = a.Clusters
	v.Platforms = a.Platforms
	v.Sites = a.Sites
	v.Tenants = a.Tenants
	v.Tags = a.Tags
}

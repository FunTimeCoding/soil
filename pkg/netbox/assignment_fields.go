package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_request"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

func (c *Client) assignmentFields(
	q *configuration_request.Request,
	a *configuration_assignment.Assignment,
) error {
	roles, e := resolveIdentifiers(
		a.Roles,
		c.DeviceRoleByName,
		func(v *device_role.Role) int32 { return v.Identifier },
	)

	if e != nil {
		return e
	}

	clusters, e := resolveIdentifiers(
		a.Clusters,
		c.ClusterByName,
		func(v *cluster.Cluster) int32 { return v.Identifier },
	)

	if e != nil {
		return e
	}

	platforms, e := resolveIdentifiers(
		a.Platforms,
		c.PlatformByName,
		func(v *platform.Platform) int32 { return v.Identifier },
	)

	if e != nil {
		return e
	}

	sites, e := resolveIdentifiers(
		a.Sites,
		c.SiteByName,
		func(v *site.Site) int32 { return v.Identifier },
	)

	if e != nil {
		return e
	}

	tenants, e := resolveIdentifiers(
		a.Tenants,
		c.TenantByName,
		func(v *tenant.Tenant) int32 { return v.Identifier },
	)

	if e != nil {
		return e
	}

	tags := []string{}

	for _, n := range a.Tags {
		t, f := c.TagByName(n)

		if f != nil {
			return f
		}

		tags = append(tags, t.Raw.GetSlug())
	}

	q.Roles = &roles
	q.Clusters = &clusters
	q.Platforms = &platforms
	q.Sites = &sites
	q.Tenants = &tenants
	q.Tags = &tags

	return nil
}

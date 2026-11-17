package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/location"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestDeviceRoleSlugReported(t *testing.T) {
	v := netbox.NewDeviceRoleWithDefaults()
	v.SetSlug("alfa-bravo")
	assert.String(t, "alfa-bravo", convert.DeviceRole(device_role.New(v)).Slug)
}

func TestClusterTypeSlugReported(t *testing.T) {
	v := netbox.NewClusterTypeWithDefaults()
	v.SetSlug("alfa-bravo")
	assert.String(
		t,
		"alfa-bravo",
		convert.ClusterType(cluster_type.New(v)).Slug,
	)
}

func TestLocationSlugReported(t *testing.T) {
	v := netbox.NewLocationWithDefaults()
	v.SetSlug("alfa-bravo")
	assert.String(t, "alfa-bravo", convert.Location(location.New(v)).Slug)
}

func TestDeviceTypeSlugReported(t *testing.T) {
	v := netbox.NewDeviceTypeWithDefaults()
	v.SetSlug("alfa-bravo")
	assert.String(t, "alfa-bravo", convert.DeviceType(device_type.New(v)).Slug)
}

func TestPlatformSlugReported(t *testing.T) {
	v := netbox.NewPlatformWithDefaults()
	v.SetSlug("alfa-bravo")
	assert.String(t, "alfa-bravo", convert.Platform(platform.New(v)).Slug)
}

func TestSiteSlugReported(t *testing.T) {
	v := netbox.NewSiteWithDefaults()
	v.SetSlug("alfa-bravo")
	assert.String(t, "alfa-bravo", convert.Site(site.New(v)).Slug)
}

func TestTagSlugReported(t *testing.T) {
	v := netbox.NewTagWithDefaults()
	v.SetSlug("alfa-bravo")
	assert.String(t, "alfa-bravo", convert.Tag(tag.New(v)).Slug)
}

func TestManufacturerSlugReported(t *testing.T) {
	v := netbox.NewManufacturerWithDefaults()
	v.SetSlug("alfa-bravo")
	assert.String(
		t,
		"alfa-bravo",
		convert.Manufacturer(manufacturer.New(v)).Slug,
	)
}

func TestTenantSlugReported(t *testing.T) {
	v := netbox.NewTenantWithDefaults()
	v.SetSlug("alfa-bravo")
	assert.String(t, "alfa-bravo", convert.Tenant(tenant.New(v)).Slug)
}

func TestTunnelGroupSlugReported(t *testing.T) {
	v := netbox.NewTunnelGroupWithDefaults()
	v.SetSlug("alfa-bravo")
	assert.String(
		t,
		"alfa-bravo",
		convert.TunnelGroup(tunnel_group.New(v)).Slug,
	)
}

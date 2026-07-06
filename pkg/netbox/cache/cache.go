package cache

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

type Cache struct {
	DeviceRoles       []*device_role.Role
	DeviceTypes       []*device_type.Type
	InternetAddresses []*internet_address.Address
	Manufacturers     []*manufacturer.Manufacturer
	PhysicalAddresses []*physical_address.Address
	Prefixes          []*prefix.Prefix
	Racks             []*rack.Rack
	Sites             []*site.Site
	Tags              []*tag.Tag
	Tenants           []*tenant.Tenant
}

package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/contact"
	"github.com/funtimecoding/soil/pkg/netbox/contact_group"
	"github.com/funtimecoding/soil/pkg/netbox/contact_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/location"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/netbox/tenant_group"
)

type OrganizationClient interface {
	CreateDeviceRole(
		name string,
		virtualMachine bool,
	) (*device_role.Role, error)
	CreateDeviceType(
		model string,
		m *manufacturer.Manufacturer,
	) (*device_type.Type, error)
	CreateLocation(
		name string,
		siteName string,
	) (*location.Location, error)
	CreateManufacturer(name string) (*manufacturer.Manufacturer, error)
	CreatePlatform(name string) (*platform.Platform, error)
	CreateSite(name string) (*site.Site, error)
	CreateTenant(name string) (*tenant.Tenant, error)
	DeviceRoleByName(n string) (*device_role.Role, error)
	DeviceRoles() ([]*device_role.Role, error)
	DeviceTypeByName(n string) (*device_type.Type, error)
	DeviceTypes() ([]*device_type.Type, error)
	Locations() ([]*location.Location, error)
	ManufacturerByName(n string) (*manufacturer.Manufacturer, error)
	Manufacturers() ([]*manufacturer.Manufacturer, error)
	Platforms() ([]*platform.Platform, error)
	SiteByName(n string) (*site.Site, error)
	Sites() ([]*site.Site, error)
	TenantByName(n string) (*tenant.Tenant, error)
	Tenants() ([]*tenant.Tenant, error)
	ContactGroups() ([]*contact_group.Group, error)
	ContactRoles() ([]*contact_role.Role, error)
	Contacts() ([]*contact.Contact, error)
	LocationByName(n string) (*location.Location, error)
	MustContactGroups() []*contact_group.Group
	MustContactRoles() []*contact_role.Role
	MustContacts() []*contact.Contact
	MustDeviceRoleByName(n string) *device_role.Role
	MustDeviceRoles() []*device_role.Role
	MustDeviceTypeByName(n string) *device_type.Type
	MustDeviceTypes() []*device_type.Type
	MustLocations() []*location.Location
	MustManufacturers() []*manufacturer.Manufacturer
	MustPlatforms() []*platform.Platform
	MustSiteByName(n string) *site.Site
	MustSites() []*site.Site
	MustTenantByName(n string) *tenant.Tenant
	MustTenantGroups() []*tenant_group.Group
	MustTenants() []*tenant.Tenant
	PlatformByName(n string) (*platform.Platform, error)
	SitesByMatch(m string) ([]*site.Site, error)
	SitesByName(n string) ([]*site.Site, error)
	TenantGroups() ([]*tenant_group.Group, error)
}

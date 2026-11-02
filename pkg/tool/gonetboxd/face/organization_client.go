package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

type OrganizationClient interface {
	CreateLocation(
		x context.Context,
		name string,
		site string,
	) (client.Location, error)
	CreateManufacturer(
		x context.Context,
		name string,
	) (client.Manufacturer, error)
	CreatePlatform(
		x context.Context,
		name string,
	) (client.Platform, error)
	CreateSite(
		x context.Context,
		name string,
	) (client.Site, error)
	CreateTenant(
		x context.Context,
		name string,
	) (client.Tenant, error)
	ListLocations(
		x context.Context,
	) ([]*client.Location, error)
	ListManufacturers(
		x context.Context,
	) ([]*client.Manufacturer, error)
	ListPlatforms(
		x context.Context,
	) ([]*client.Platform, error)
	ListSites(
		x context.Context,
	) ([]*client.Site, error)
	ListTenants(
		x context.Context,
	) ([]*client.Tenant, error)
}

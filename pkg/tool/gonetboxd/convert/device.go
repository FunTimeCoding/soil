package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Device(d *device.Device) *server.Device {
	result := &server.Device{
		Identifier: d.Identifier,
		Name:       d.Name,
	}

	if d.Role != "" {
		result.Role = &d.Role
	}

	if d.Type != "" {
		result.Type = &d.Type
	}

	if d.Site != "" {
		result.Site = &d.Site
	}

	if d.Tenant != "" {
		result.Tenant = &d.Tenant
	}

	if d.Serial != "" {
		result.Serial = &d.Serial
	}

	if d.PrimaryAddress != "" {
		result.PrimaryAddress = &d.PrimaryAddress
	}

	if len(d.Tags) > 0 {
		result.Tags = &d.Tags
	}

	if d.Link != "" {
		result.Link = &d.Link
	}

	return result
}

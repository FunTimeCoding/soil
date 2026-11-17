package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func DeviceType(t *device_type.Type) *server.DeviceType {
	result := &server.DeviceType{
		Identifier: t.Identifier,
		Model:      t.Model,
		Slug:       t.Slug,
	}

	if t.Raw.Manufacturer.Name != "" {
		result.Manufacturer = &t.Raw.Manufacturer.Name
	}

	return result
}

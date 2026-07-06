package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func DeviceRole(r *device_role.Role) *server.DeviceRole {
	return &server.DeviceRole{Identifier: r.Identifier, Name: r.Name}
}

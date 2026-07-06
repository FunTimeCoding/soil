package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func DeviceRoles(v []*device_role.Role) []*server.DeviceRole {
	result := make([]*server.DeviceRole, 0, len(v))

	for _, r := range v {
		result = append(result, DeviceRole(r))
	}

	return result
}

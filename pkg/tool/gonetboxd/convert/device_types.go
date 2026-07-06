package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func DeviceTypes(v []*device_type.Type) []*server.DeviceType {
	result := make([]*server.DeviceType, 0, len(v))

	for _, t := range v {
		result = append(result, DeviceType(t))
	}

	return result
}

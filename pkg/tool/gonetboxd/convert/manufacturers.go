package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Manufacturers(v []*manufacturer.Manufacturer) []*server.Manufacturer {
	result := make([]*server.Manufacturer, 0, len(v))

	for _, m := range v {
		result = append(result, Manufacturer(m))
	}

	return result
}

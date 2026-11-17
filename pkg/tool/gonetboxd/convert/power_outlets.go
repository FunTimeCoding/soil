package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func PowerOutlets(v []*power_outlet.Outlet) []*server.PowerOutlet {
	result := make([]*server.PowerOutlet, 0, len(v))

	for _, o := range v {
		result = append(result, PowerOutlet(o))
	}

	return result
}

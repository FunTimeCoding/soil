package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Cables(v []*cable.Cable) []*server.Cable {
	result := make([]*server.Cable, 0, len(v))

	for _, c := range v {
		result = append(result, Cable(c))
	}

	return result
}

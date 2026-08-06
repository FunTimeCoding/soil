package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/location"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Locations(v []*location.Location) []*server.Location {
	result := make([]*server.Location, 0, len(v))

	for _, l := range v {
		result = append(result, Location(l))
	}

	return result
}

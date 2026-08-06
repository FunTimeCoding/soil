package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Platforms(v []*platform.Platform) []*server.Platform {
	result := make([]*server.Platform, 0, len(v))

	for _, p := range v {
		result = append(result, Platform(p))
	}

	return result
}

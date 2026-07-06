package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Prefixes(v []*prefix.Prefix) []*server.Prefix {
	result := make([]*server.Prefix, 0, len(v))

	for _, p := range v {
		result = append(result, Prefix(p))
	}

	return result
}

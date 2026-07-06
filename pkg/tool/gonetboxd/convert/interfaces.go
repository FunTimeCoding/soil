package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Interfaces(v []*network.Interface) []*server.Interface {
	result := make([]*server.Interface, 0, len(v))

	for _, i := range v {
		result = append(result, Interface(i))
	}

	return result
}

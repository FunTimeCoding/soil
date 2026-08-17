package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/network_interface"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func NetworkInterfaces(v []*network_interface.Interface) []server.NetworkInterface {
	result := []server.NetworkInterface{}

	for _, e := range v {
		result = append(result, *NetworkInterface(e))
	}

	return result
}

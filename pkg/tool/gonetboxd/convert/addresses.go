package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Addresses(v []*internet_address.Address) []*server.Address {
	result := make([]*server.Address, 0, len(v))

	for _, a := range v {
		result = append(result, Address(a))
	}

	return result
}

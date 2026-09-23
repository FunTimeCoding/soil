package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func PhysicalAddressOwners(
	v []*physical_address.Address,
) []server.PhysicalAddressOwner {
	result := make([]server.PhysicalAddressOwner, 0, len(v))

	for _, e := range v {
		result = append(result, *PhysicalAddressOwner(e))
	}

	return result
}

package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Interface(i *network.Interface) *server.Interface {
	result := &server.Interface{Identifier: i.Identifier, Name: i.Name}

	if i.Description != "" {
		result.Description = &i.Description
	}

	if i.Type != "" {
		result.Type = new(string(i.Type))
	}

	if len(i.PhysicalAddress) > 0 {
		result.PhysicalAddress = new(i.PhysicalAddress.String())
	}

	return result
}

package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Cable(c *cable.Cable) *server.Cable {
	result := &server.Cable{Identifier: c.Identifier, Name: c.Name}

	if c.Status != "" {
		result.Status = &c.Status
	}

	if c.SideA != "" {
		result.SideA = &c.SideA
	}

	if c.SideB != "" {
		result.SideB = &c.SideB
	}

	if c.Link != "" {
		result.Link = &c.Link
	}

	return result
}

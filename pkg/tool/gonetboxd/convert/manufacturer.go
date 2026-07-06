package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Manufacturer(m *manufacturer.Manufacturer) *server.Manufacturer {
	return &server.Manufacturer{Identifier: m.Identifier, Name: m.Name}
}

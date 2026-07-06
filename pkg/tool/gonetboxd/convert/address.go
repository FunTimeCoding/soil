package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Address(a *internet_address.Address) *server.Address {
	result := &server.Address{
		Identifier: a.Identifier,
		Address:    a.Address.String(),
	}

	if a.Name != "" {
		result.Name = &a.Name
	}

	if a.ObjectType != "" {
		result.ObjectType = &a.ObjectType
	}

	return result
}

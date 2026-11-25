package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Address(a *internet_address.Address) *server.Address {
	return &server.Address{
		Identifier:     a.Identifier,
		Address:        a.Address.String(),
		Name:           filled(a.Name),
		ObjectType:     filled(a.ObjectType),
		Interface:      filled(a.Interface),
		Device:         filled(a.Device),
		VirtualMachine: filled(a.VirtualMachine),
		Role:           filled(a.Role),
		Status:         filled(a.Status),
		DomainName:     filled(a.DomainName),
		Description:    filled(a.Description),
	}
}

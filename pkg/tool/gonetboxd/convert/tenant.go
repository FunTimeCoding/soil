package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Tenant(t *tenant.Tenant) *server.Tenant {
	return &server.Tenant{Identifier: t.Identifier, Name: t.Name}
}

package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Tenants(v []*tenant.Tenant) []*server.Tenant {
	result := make([]*server.Tenant, 0, len(v))

	for _, t := range v {
		result = append(result, Tenant(t))
	}

	return result
}

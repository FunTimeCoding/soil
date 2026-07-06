package tenant

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/netbox-community/go-netbox/v4"
)

func New(t *netbox.Tenant) *Tenant {
	result := &Tenant{Identifier: t.GetId(), Name: t.GetName(), Raw: t}

	if t.Group.IsSet() {
		result.Group = t.Group.Get().GetName()
	} else {
		result.Group = constant.NoGroup
	}

	return result
}

package virtual_machine

import (
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/netbox-community/go-netbox/v4"
)

func New(v *netbox.VirtualMachineWithConfigContext) *Machine {
	return &Machine{
		Identifier: v.GetId(),
		Name:       v.GetName(),
		Tags:       tag.Names(v.Tags),
		Raw:        v,
	}
}

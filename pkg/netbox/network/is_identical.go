package network

import "github.com/funtimecoding/soil/pkg/netbox/types/network_definition"

func (i *Interface) IsIdentical(d *network_definition.Definition) bool {
	if string(i.Type) != d.Type {
		return false
	}

	if i.PhysicalAddress.String() != d.PhysicalAddress.String() {
		return false
	}

	return true
}

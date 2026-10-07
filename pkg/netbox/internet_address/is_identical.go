package internet_address

import "github.com/funtimecoding/soil/pkg/netbox/types/internet_address_definition"

func (a *Address) IsIdentical(d *internet_address_definition.Definition) bool {
	if !a.Address.Equal(d.Address) {
		return false
	}

	if a.Network.Mask.String() != d.Mask.String() {
		return false
	}

	return true
}

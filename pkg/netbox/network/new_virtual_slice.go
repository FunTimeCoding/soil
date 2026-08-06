package network

import "github.com/netbox-community/go-netbox/v4"

func NewVirtualSlice(v []netbox.VMInterface) []*Interface {
	var result []*Interface

	for _, e := range v {
		result = append(result, NewVirtual(&e))
	}

	return result
}

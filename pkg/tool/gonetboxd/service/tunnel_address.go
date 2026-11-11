package service

import "github.com/funtimecoding/soil/pkg/netbox/internet_address"

func tunnelAddress(
	ends map[string]bool,
	addresses []*internet_address.Address,
) string {
	for _, a := range addresses {
		if ends[terminationKey(a.ObjectType, a.ObjectIdentifier)] {
			return a.Address.String()
		}
	}

	return ""
}

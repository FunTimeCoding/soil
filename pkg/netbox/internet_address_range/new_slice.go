package internet_address_range

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.IPRange) []*Range {
	var result []*Range

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}

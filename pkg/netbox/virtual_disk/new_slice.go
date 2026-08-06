package virtual_disk

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.VirtualDisk) []*Disk {
	var result []*Disk

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}

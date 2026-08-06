package virtual_disk

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.VirtualDisk) *Disk {
	return &Disk{
		Identifier: v.GetId(),
		Name:       v.GetName(),
		Size:       v.GetSize(),
		Raw:        v,
	}
}

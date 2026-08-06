package virtual_disk

import "github.com/netbox-community/go-netbox/v4"

type Disk struct {
	Identifier int32
	Name       string
	Size       int32
	Raw        *netbox.VirtualDisk
}

package manufacturer

import "github.com/netbox-community/go-netbox/v4"

type Manufacturer struct {
	Identifier int32
	Name       string
	Slug       string
	Raw        *netbox.Manufacturer
}

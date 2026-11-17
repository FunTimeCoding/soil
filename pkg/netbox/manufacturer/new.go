package manufacturer

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.Manufacturer) *Manufacturer {
	return &Manufacturer{
		Identifier: v.GetId(),
		Name:       v.GetName(),
		Slug:       v.GetSlug(),
		Raw:        v,
	}
}

package location

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.Location) *Location {
	return &Location{
		Identifier: v.GetId(),
		Name:       v.GetDisplay(),
		Slug:       v.GetSlug(),
		Raw:        v,
	}
}

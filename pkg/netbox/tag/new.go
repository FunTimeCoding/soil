package tag

import "github.com/netbox-community/go-netbox/v4"

func New(t *netbox.Tag) *Tag {
	return &Tag{
		Identifier: t.GetId(),
		Name:       t.GetName(),
		Slug:       t.GetSlug(),
		Raw:        t,
		Nested: &netbox.NestedTagRequest{
			Color: t.Color,
			Name:  t.Name,
			Slug:  t.Slug,
		},
	}
}

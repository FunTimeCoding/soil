package site

import "github.com/netbox-community/go-netbox/v4"

func New(s *netbox.Site) *Site {
	return &Site{
		Identifier: s.GetId(),
		Name:       s.GetName(),
		Slug:       s.GetSlug(),
		Raw:        s,
	}
}

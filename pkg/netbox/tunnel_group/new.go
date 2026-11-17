package tunnel_group

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.TunnelGroup) *Group {
	return &Group{
		Identifier: v.GetId(),
		Name:       v.GetName(),
		Slug:       v.GetSlug(),
		Raw:        v,
	}
}

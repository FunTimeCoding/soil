package cluster_type

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.ClusterType) *Type {
	return &Type{
		Identifier: v.GetId(),
		Name:       v.GetName(),
		Slug:       v.GetSlug(),
		Raw:        v,
	}
}

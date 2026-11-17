package device_type

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.DeviceType) *Type {
	return &Type{
		Identifier: v.GetId(),
		Model:      v.GetModel(),
		Slug:       v.GetSlug(),
		Raw:        v,
	}
}

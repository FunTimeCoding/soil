package device

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/helper"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/netbox-community/go-netbox/v4"
)

func New(v *netbox.DeviceWithConfigContext) *Device {
	var address string
	var name string
	var tenant string
	var comment string

	if v.PrimaryIp4.IsSet() {
		address = v.PrimaryIp4.Get().GetDisplay()
	} else {
		address = constant.NoPrimaryAddress
	}

	if v.Name.IsSet() {
		name = v.GetName()
	} else {
		name = constant.NoName
	}

	if v.Tenant.IsSet() {
		tenant = v.Tenant.Get().GetName()
	} else {
		tenant = constant.NoTenant
	}

	if v.Comments == nil {
		comment = constant.NoComment
	} else {
		comment = v.GetComments()
	}

	return &Device{
		Identifier:     v.GetId(),
		Name:           name,
		Role:           v.Role.GetName(),
		Type:           v.DeviceType.GetDisplay(),
		Site:           v.Site.GetName(),
		Tenant:         tenant,
		Serial:         v.GetSerial(),
		Comment:        comment,
		PrimaryAddress: address,
		Tags:           tag.Names(v.Tags),
		Link:           helper.ToWebLink(v.GetUrl()),
		Raw:            v,
	}
}

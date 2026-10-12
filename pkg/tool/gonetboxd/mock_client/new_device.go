package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/netbox-community/go-netbox/v4"
)

func newDevice(
	identifier int32,
	name string,
	tags []string,
	n *tenant.Tenant,
) *device.Device {
	v := netbox.NewDeviceWithConfigContextWithDefaults()
	v.SetId(identifier)
	v.SetName(name)

	for _, t := range tags {
		v.Tags = append(v.Tags, *netbox.NewNestedTag(0, "", t, t, t))
	}

	if n != nil {
		v.SetTenant(
			*netbox.NewBriefTenant(n.Identifier, "", n.Name, n.Name, n.Name),
		)
	}

	return device.New(v)
}

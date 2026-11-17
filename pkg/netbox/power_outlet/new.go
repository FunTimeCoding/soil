package power_outlet

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.PowerOutlet) *Outlet {
	result := &Outlet{
		Identifier: v.GetId(),
		Name:       v.GetName(),
		Device:     v.Device.GetName(),
		Raw:        v,
	}
	result.ConnectedDevice, result.ConnectedPort = peer(v.LinkPeers)

	return result
}

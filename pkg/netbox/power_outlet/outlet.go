package power_outlet

import "github.com/netbox-community/go-netbox/v4"

type Outlet struct {
	Identifier      int32
	Name            string
	Device          string
	ConnectedDevice string
	ConnectedPort   string
	Raw             *netbox.PowerOutlet
}

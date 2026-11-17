package device_type

import "github.com/netbox-community/go-netbox/v4"

type Type struct {
	Identifier int32
	Model      string
	Slug       string
	Raw        *netbox.DeviceType
}

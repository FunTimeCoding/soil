package physical_address

import (
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

type Address struct {
	Identifier       int32
	Name             string
	Address          net.HardwareAddr
	ObjectType       string
	ObjectIdentifier int64
	Interface        *netbox.BriefInterface
	VirtualInterface *VirtualInterface
	Raw              *netbox.MACAddress
}

package internet_address

import (
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

type Address struct {
	Identifier       int32
	Name             string
	Address          net.IP
	Network          *net.IPNet
	ObjectType       string
	ObjectIdentifier int64
	Interface        string
	Device           string
	VirtualMachine   string
	Role             string
	Status           string
	DomainName       string
	Description      string
	Raw              *netbox.IPAddress
}

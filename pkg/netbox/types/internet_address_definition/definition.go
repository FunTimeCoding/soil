package internet_address_definition

import "net"

type Definition struct {
	Name    string
	Address net.IP
	Mask    net.IPMask
}

package network_definition

import "net"

type Definition struct {
	Name            string
	Type            string
	PhysicalAddress net.HardwareAddr
}

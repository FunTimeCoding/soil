package virtual_machine

import "github.com/netbox-community/go-netbox/v4"

type Machine struct {
	Identifier     int32
	Name           string
	Cluster        string
	Site           string
	Status         string
	Role           string
	PrimaryAddress string
	Tags           []string
	Configuration  map[string]any
	CustomFields   map[string]any
	Raw            *netbox.VirtualMachineWithConfigContext
}

package virtual_interface

import "github.com/netbox-community/go-netbox/v4"

type Interface struct {
	Identifier     int32                      `json:"id"`
	Name           string                     `json:"name"`
	VirtualMachine netbox.BriefVirtualMachine `json:"virtual_machine"`
}

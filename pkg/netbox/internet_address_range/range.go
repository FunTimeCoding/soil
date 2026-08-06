package internet_address_range

import "github.com/netbox-community/go-netbox/v4"

type Range struct {
	Identifier int32
	Start      string
	End        string
	Status     string
	Raw        *netbox.IPRange
}

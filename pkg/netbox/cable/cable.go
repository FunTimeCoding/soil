package cable

import "github.com/netbox-community/go-netbox/v4"

type Cable struct {
	Identifier  int32
	Name        string
	Description string
	Status      string
	SideA       string
	SideB       string
	Link        string
	Raw         *netbox.Cable
}

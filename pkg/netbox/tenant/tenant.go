package tenant

import "github.com/netbox-community/go-netbox/v4"

type Tenant struct {
	Identifier int32
	Group      string
	Name       string
	Slug       string
	Raw        *netbox.Tenant
}

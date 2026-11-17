package tunnel_group

import "github.com/netbox-community/go-netbox/v4"

type Group struct {
	Identifier int32
	Name       string
	Slug       string
	Raw        *netbox.TunnelGroup
}

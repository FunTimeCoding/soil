package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func WirelessNetwork(n *wireless_network.Network) *server.WirelessNetwork {
	return &server.WirelessNetwork{Identifier: n.Identifier, Name: n.Name}
}

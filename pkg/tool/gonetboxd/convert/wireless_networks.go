package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func WirelessNetworks(v []*wireless_network.Network) []*server.WirelessNetwork {
	result := make([]*server.WirelessNetwork, 0, len(v))

	for _, n := range v {
		result = append(result, WirelessNetwork(n))
	}

	return result
}

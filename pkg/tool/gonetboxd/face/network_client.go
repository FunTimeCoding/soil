package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

type NetworkClient interface {
	CreateDeviceTunnelTermination(
		x context.Context,
		device string,
		tunnel string,
		interfaceName string,
		role string,
	) (client.TunnelTermination, error)
	CreateTunnel(
		x context.Context,
		name string,
		encapsulation string,
		group string,
	) (client.Tunnel, error)
	CreateTunnelGroup(
		x context.Context,
		name string,
	) (client.TunnelGroup, error)
	CreateWirelessNetwork(
		x context.Context,
		ssid string,
	) (client.WirelessNetwork, error)
	ListTunnelGroups(
		x context.Context,
	) ([]*client.TunnelGroup, error)
	ListTunnelTerminations(
		x context.Context,
	) ([]*client.TunnelTermination, error)
	ListTunnels(
		x context.Context,
	) ([]*client.Tunnel, error)
	ListWirelessNetworks(
		x context.Context,
	) ([]*client.WirelessNetwork, error)
}

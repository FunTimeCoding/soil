package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_network"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_network_group"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_link"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network_group"
)

type NetworkClient interface {
	CreateTunnel(
		name string,
		encapsulation string,
		group *tunnel_group.Group,
	) (*tunnel.Tunnel, error)
	CreateTunnelGroup(name string) (*tunnel_group.Group, error)
	CreateTunnelTermination(
		t *tunnel.Tunnel,
		terminationType string,
		terminationIdentifier int64,
		role string,
	) (*tunnel_termination.Termination, error)
	CreateWirelessNetwork(ssid string) (*wireless_network.Network, error)
	TunnelByName(n string) (*tunnel.Tunnel, error)
	TunnelGroupByName(n string) (*tunnel_group.Group, error)
	TunnelGroups() ([]*tunnel_group.Group, error)
	Tunnels() ([]*tunnel.Tunnel, error)
	TunnelTerminations() ([]*tunnel_termination.Termination, error)
	WirelessNetworks() ([]*wireless_network.Network, error)
	MustTunnelGroups() []*tunnel_group.Group
	MustTunnels() []*tunnel.Tunnel
	MustVirtualNetworkGroups() []*virtual_network_group.Group
	MustVirtualNetworks() []*virtual_network.Network
	MustWirelessLinks() []*wireless_link.Link
	MustWirelessNetworkGroups() []*wireless_network_group.Group
	MustWirelessNetworks() []*wireless_network.Network
	VirtualNetworkGroups() ([]*virtual_network_group.Group, error)
	VirtualNetworks() ([]*virtual_network.Network, error)
	WirelessLinks() ([]*wireless_link.Link, error)
	WirelessNetworkGroups() ([]*wireless_network_group.Group, error)
}

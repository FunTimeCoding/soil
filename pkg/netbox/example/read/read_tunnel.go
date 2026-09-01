package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readTunnel(
	n *netbox.Client,
	f *option.Format,
) {
	// VPN
	for _, g := range n.MustTunnelGroups() {
		console.Format("TunnelGroup: %s\n", g.Format(f))
	}

	for _, t := range n.MustTunnels() {
		console.Format("Tunnel: %s\n", t.Format(f))
	}
}

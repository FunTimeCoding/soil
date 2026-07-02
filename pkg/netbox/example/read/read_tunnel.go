package read

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func readTunnel(
	n *netbox.Client,
	f *option.Format,
) {
	// VPN
	for _, g := range n.MustTunnelGroups() {
		fmt.Printf("TunnelGroup: %s\n", g.Format(f))
	}

	for _, t := range n.MustTunnels() {
		fmt.Printf("Tunnel: %s\n", t.Format(f))
	}
}

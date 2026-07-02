package read

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func readWireless(
	n *netbox.Client,
	f *option.Format,
) {
	for _, g := range n.MustWirelessNetworkGroups() {
		fmt.Printf("WirelessNetworkGroup: %s\n", g.Format(f))
	}

	for _, e := range n.MustWirelessNetworks() {
		fmt.Printf("WirelessNetwork: %s\n", e.Format(f))
	}

	if false {
		// TODO: What must devices have to show up in the picker?
		for _, l := range n.MustWirelessLinks() {
			fmt.Printf("WirelessLink: %s\n", l.Format(f))
		}
	}
}

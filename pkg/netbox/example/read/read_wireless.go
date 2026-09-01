package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readWireless(
	n *netbox.Client,
	f *option.Format,
) {
	for _, g := range n.MustWirelessNetworkGroups() {
		console.Format("WirelessNetworkGroup: %s\n", g.Format(f))
	}

	for _, e := range n.MustWirelessNetworks() {
		console.Format("WirelessNetwork: %s\n", e.Format(f))
	}

	if false {
		// TODO: What must devices have to show up in the picker?
		for _, l := range n.MustWirelessLinks() {
			console.Format("WirelessLink: %s\n", l.Format(f))
		}
	}
}

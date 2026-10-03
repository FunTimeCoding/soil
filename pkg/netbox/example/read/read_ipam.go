package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readIPAM(
	n *netbox.Client,
	f *option.Format,
) {
	for _, t := range n.MustServiceTemplates() {
		console.Format("ServiceTemplate: %s\n", t.Format(f))
	}

	for _, s := range n.MustServices() {
		console.Format("Services: %s\n", s.Format(f))
	}

	if false {
		for _, g := range n.MustVirtualNetworkGroups() {
			console.Format("VirtualNetworkGroup: %s\n", g.Format(f))
		}
	}

	for _, e := range n.MustVirtualNetworks() {
		console.Format("VirtualNetwork: %s\n", e.Format(f))
	}

	for _, u := range n.MustSystemNumbers() {
		console.Format("SystemNumber: %s\n", u.Format(f))
	}
}

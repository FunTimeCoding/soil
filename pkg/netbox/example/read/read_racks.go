package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readRacks(n *netbox.Client, f *option.Format) {
	for _, t := range n.MustRackTypes() {
		console.Format("RackType: %s\n", t.Format(f))
	}

	for _, r := range n.MustRackRoles() {
		console.Format("RackRole: %s\n", r.Format(f))
	}

	for _, r := range n.MustRacks() {
		console.Format("Rack: %s\n", r.Format(f))
	}

	for _, p := range n.MustCables() {
		console.Format("Cable: %s\n", p.Format(f))
	}
}

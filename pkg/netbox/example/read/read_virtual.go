package read

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readVirtual(
	n *netbox.Client,
	f *option.Format,
) {
	// Virtualization
	for _, g := range n.MustClusterGroups() {
		console.Format("ClusterGroup: %s\n", g.Format(f))
	}

	for _, t := range n.MustClusterTypes() {
		console.Format("ClusterType: %s\n", t.Format(f))
	}

	for _, c := range n.MustClusters() {
		console.Format("Cluster: %s\n", c.Format(f))
	}

	for _, m := range n.MustVirtualMachines() {
		console.Format("VirtualMachine: %s\n", m.Format(f))
	}
}

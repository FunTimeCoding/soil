package read

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func readVirtual(
	n *netbox.Client,
	f *option.Format,
) {
	// Virtualization
	for _, g := range n.MustClusterGroups() {
		fmt.Printf("ClusterGroup: %s\n", g.Format(f))
	}

	for _, t := range n.MustClusterTypes() {
		fmt.Printf("ClusterType: %s\n", t.Format(f))
	}

	for _, c := range n.MustClusters() {
		fmt.Printf("Cluster: %s\n", c.Format(f))
	}

	for _, m := range n.MustVirtualMachines() {
		fmt.Printf("VirtualMachine: %s\n", m.Format(f))
	}
}

package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/network"
)

func Physical() {
	n := netbox.NewEnvironment()
	defer n.Close()

	for _, p := range n.MustHardwarePhysicalAddresses(
		network.PhysicalAddress(soil.PhysicalTest0),
	) {
		console.Format("Read physical address: %+v\n", p)
	}
}

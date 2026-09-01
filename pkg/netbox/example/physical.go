package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/network"
)

func Physical() {
	n := netbox.NewEnvironment()

	for _, p := range n.MustPhysicalAddressesByHardware(
		network.PhysicalAddress(constant.PhysicalTest0),
	) {
		console.Format("Read physical address: %+v\n", p)
	}
}

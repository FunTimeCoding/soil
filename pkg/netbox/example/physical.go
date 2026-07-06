package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/network"
)

func Physical() {
	n := netbox.NewEnvironment()

	for _, p := range n.MustPhysicalAddressesByHardware(
		network.PhysicalAddress(constant.PhysicalTest0),
	) {
		fmt.Printf("Read physical address: %+v\n", p)
	}
}

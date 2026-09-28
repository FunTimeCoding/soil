package gazetteer

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func New(
	devices []*client.Device,
	machines []*client.VirtualMachine,
	physical []client.PhysicalAddressOwner,
) *Gazetteer {
	result := &Gazetteer{
		place:    map[string]*place.Place{},
		address:  map[string]*place.Place{},
		hardware: map[string]*place.Place{},
	}

	for _, m := range machines {
		result.add(
			place.New(constant.VirtualMachineAddress, m.Identifier, m.Name),
			m.PrimaryAddress,
		)
	}

	for _, d := range devices {
		result.add(
			place.New(constant.DeviceAddress, d.Identifier, d.Name),
			d.PrimaryAddress,
		)
	}

	for _, p := range physical {
		result.addHardware(&p)
	}

	return result
}

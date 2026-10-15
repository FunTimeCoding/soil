package ansible_inventory

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func New(
	field string,
	machines []*virtual_machine.Machine,
	devices []*device.Device,
) *Inventory {
	result := &Inventory{groups: map[string]map[string]string{}}

	for _, m := range machines {
		for _, g := range selected(m.CustomFields, field) {
			result.add(g, m.Name, hostAddress(m.PrimaryAddress))
		}
	}

	for _, d := range devices {
		for _, g := range selected(d.CustomFields, field) {
			result.add(g, d.Name, hostAddress(d.PrimaryAddress))
		}
	}

	return result
}

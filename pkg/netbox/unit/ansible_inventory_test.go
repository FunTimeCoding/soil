package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/ansible_inventory"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestAnsibleInventoryGroupsHostsByField(t *testing.T) {
	m := netbox.NewVirtualMachineWithConfigContextWithDefaults()
	m.SetName("alfa")
	m.SetCustomFields(
		map[string]any{"ansible_playbook": []any{"echo", "foxtrot"}},
	)
	m.SetPrimaryIp4(
		*netbox.NewBriefIPAddress(
			1,
			"",
			"192.168.0.5/24",
			netbox.AggregateFamily{},
			"192.168.0.5/24",
		),
	)
	d := netbox.NewDeviceWithConfigContextWithDefaults()
	d.SetName("bravo")
	d.SetCustomFields(map[string]any{"ansible_playbook": []any{"echo"}})
	unmanaged := netbox.NewVirtualMachineWithConfigContextWithDefaults()
	unmanaged.SetName("charlie")
	assert.String(
		t,
		`all:
    children:
        echo:
            hosts:
                alfa:
                    ansible_host: 192.168.0.5
                bravo: {}
        foxtrot:
            hosts:
                alfa:
                    ansible_host: 192.168.0.5
`,
		ansible_inventory.New(
			"ansible_playbook",
			[]*virtual_machine.Machine{
				virtual_machine.New(m),
				virtual_machine.New(unmanaged),
			},
			[]*device.Device{device.New(d)},
		).Markup(),
	)
}

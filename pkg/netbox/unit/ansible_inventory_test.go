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

func TestAnsibleInventoryCarriesContextVariables(t *testing.T) {
	m := netbox.NewVirtualMachineWithConfigContextWithDefaults()
	m.SetName("alfa")
	m.SetCustomFields(map[string]any{"ansible_playbook": []any{"echo"}})
	m.SetPrimaryIp4(
		*netbox.NewBriefIPAddress(
			1,
			"",
			"192.168.0.5/24",
			netbox.AggregateFamily{},
			"192.168.0.5/24",
		),
	)
	m.SetConfigContext(
		map[string]any{
			"ansible": map[string]any{
				"ansible_port": 2200,
				"listen_port":  8080,
			},
			"backup": map[string]any{"ansible_port": 22},
		},
	)
	d := netbox.NewDeviceWithConfigContextWithDefaults()
	d.SetName("bravo")
	d.SetCustomFields(map[string]any{"ansible_playbook": []any{"echo"}})
	d.SetPrimaryIp4(
		*netbox.NewBriefIPAddress(
			2,
			"",
			"192.168.0.6/24",
			netbox.AggregateFamily{},
			"192.168.0.6/24",
		),
	)
	d.SetConfigContext(
		map[string]any{
			"ansible": map[string]any{"ansible_host": "bravo.example"},
		},
	)
	assert.String(
		t,
		`all:
    children:
        echo:
            hosts:
                alfa:
                    ansible_host: 192.168.0.5
                    ansible_port: 2200
                    listen_port: 8080
                bravo:
                    ansible_host: bravo.example
`,
		ansible_inventory.New(
			"ansible_playbook",
			[]*virtual_machine.Machine{virtual_machine.New(m)},
			[]*device.Device{device.New(d)},
		).Markup(),
	)
}

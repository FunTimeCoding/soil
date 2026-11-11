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
			map[string]string{},
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
			map[string]string{},
			[]*virtual_machine.Machine{virtual_machine.New(m)},
			[]*device.Device{device.New(d)},
		).Markup(),
	)
}

func TestAnsibleInventoryDialsTunnelAddress(t *testing.T) {
	d := netbox.NewDeviceWithConfigContextWithDefaults()
	d.SetName("alfa")
	d.SetCustomFields(map[string]any{"ansible_playbook": []any{"echo"}})
	d.SetPrimaryIp4(
		*netbox.NewBriefIPAddress(
			1,
			"",
			"203.0.113.5/32",
			netbox.AggregateFamily{},
			"203.0.113.5/32",
		),
	)
	e := netbox.NewDeviceWithConfigContextWithDefaults()
	e.SetName("bravo")
	e.SetCustomFields(map[string]any{"ansible_playbook": []any{"echo"}})
	e.SetPrimaryIp4(
		*netbox.NewBriefIPAddress(
			2,
			"",
			"192.168.0.6/24",
			netbox.AggregateFamily{},
			"192.168.0.6/24",
		),
	)
	assert.String(
		t,
		`all:
    children:
        echo:
            hosts:
                alfa:
                    ansible_host: 198.51.100.1
                bravo:
                    ansible_host: 192.168.0.6
`,
		ansible_inventory.New(
			"ansible_playbook",
			map[string]string{"alfa": "198.51.100.1"},
			[]*virtual_machine.Machine{},
			[]*device.Device{device.New(d), device.New(e)},
		).Markup(),
	)
}

func TestAnsibleInventoryContextOverridesTunnelAddress(t *testing.T) {
	d := netbox.NewDeviceWithConfigContextWithDefaults()
	d.SetName("alfa")
	d.SetCustomFields(map[string]any{"ansible_playbook": []any{"echo"}})
	d.SetConfigContext(
		map[string]any{
			"ansible": map[string]any{"ansible_host": "alfa.example"},
		},
	)
	assert.String(
		t,
		`all:
    children:
        echo:
            hosts:
                alfa:
                    ansible_host: alfa.example
`,
		ansible_inventory.New(
			"ansible_playbook",
			map[string]string{"alfa": "198.51.100.1"},
			[]*virtual_machine.Machine{},
			[]*device.Device{device.New(d)},
		).Markup(),
	)
}

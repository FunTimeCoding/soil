package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestVirtualMachineCarriesItsRole(t *testing.T) {
	v := netbox.NewVirtualMachineWithConfigContextWithDefaults()
	v.SetName("alfa")
	v.SetRole(*netbox.NewBriefDeviceRole(1, "", "", "bravo", "bravo", 0))
	assert.String(t, "bravo", virtual_machine.New(v).Role)
}

func TestVirtualMachineWithoutRole(t *testing.T) {
	v := netbox.NewVirtualMachineWithConfigContextWithDefaults()
	v.SetName("alfa")
	assert.String(t, "", virtual_machine.New(v).Role)
}

func TestVirtualMachineCarriesResolvedConfiguration(t *testing.T) {
	v := netbox.NewVirtualMachineWithConfigContextWithDefaults()
	v.SetName("alfa")
	v.SetConfigContext(map[string]any{"bravo": "charlie"})
	assert.Any(
		t,
		map[string]any{"bravo": "charlie"},
		virtual_machine.New(v).Configuration,
	)
}

func TestDeviceRoleCarriesVirtualMachineFlag(t *testing.T) {
	v := netbox.NewDeviceRoleWithDefaults()
	v.SetName("alfa")
	v.SetVmRole(true)
	assert.True(t, device_role.New(v).VirtualMachine)
}

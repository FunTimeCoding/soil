package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"slices"
	"testing"
)

func TestPhysicalAddressTargetsPermitInterfaces(t *testing.T) {
	assert.True(
		t,
		slices.Contains(
			constant.PhysicalAddressTargets,
			constant.InterfaceAddress,
		),
	)
}

func TestPhysicalAddressTargetsPermitVirtualInterfaces(t *testing.T) {
	assert.True(
		t,
		slices.Contains(
			constant.PhysicalAddressTargets,
			constant.VirtualInterfaceAddress,
		),
	)
}

func TestPhysicalAddressTargetsRefuseDevices(t *testing.T) {
	assert.False(
		t,
		slices.Contains(constant.PhysicalAddressTargets, constant.DeviceAddress),
	)
}

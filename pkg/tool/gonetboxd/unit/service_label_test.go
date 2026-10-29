package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"testing"
)

func TestDeviceCarriesItsLabels(t *testing.T) {
	c := service_tester.NewSource(t)
	c.SeedDevice("core")
	s := service_tester.Service(t, c)
	_, e := s.SetDeviceLabel("core", "rack", "a1")
	assert.FatalOnError(t, e)
	d, labels, f := s.Device("core")
	assert.FatalOnError(t, f)
	assert.String(t, "core", d.Name)
	assert.Count(t, 1, labels)
	assert.String(t, "rack", labels[0].Key)
	assert.String(t, "a1", labels[0].Value)
}

func TestDevicesGroupLabelsByDevice(t *testing.T) {
	c := service_tester.NewSource(t)
	core := c.SeedDevice("core")
	edge := c.SeedDevice("edge")
	s := service_tester.Service(t, c)
	_, e := s.SetDeviceLabel("edge", "rack", "b2")
	assert.FatalOnError(t, e)
	devices, grouped, f := s.Devices("")
	assert.FatalOnError(t, f)
	assert.Count(t, 2, devices)
	assert.Count(t, 0, grouped[core.Identifier])
	assert.Count(t, 1, grouped[edge.Identifier])
	assert.String(t, "b2", grouped[edge.Identifier][0].Value)
}

func TestDevicesQueryNarrows(t *testing.T) {
	c := service_tester.NewSource(t)
	c.SeedDevice("core")
	c.SeedDevice("edge")
	devices, _, e := service_tester.Service(t, c).Devices("ed")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, devices)
	assert.String(t, "edge", devices[0].Name)
}

func TestRemoveDeviceLabel(t *testing.T) {
	c := service_tester.NewSource(t)
	c.SeedDevice("core")
	s := service_tester.Service(t, c)
	_, e := s.SetDeviceLabel("core", "rack", "a1")
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, s.RemoveDeviceLabel("core", "rack"))
	labels, f := s.DeviceLabels("core")
	assert.FatalOnError(t, f)
	assert.Count(t, 0, labels)
}

func TestDeviceLabelRefusesUnknownDevice(t *testing.T) {
	_, e := service_tester.Service(t, service_tester.NewSource(t)).SetDeviceLabel(
		"missing",
		"rack",
		"a1",
	)
	assert.True(t, not_found.Is(e))
	assert.String(t, "device not found: missing", e.Error())
}

func TestVirtualMachineCarriesItsLabels(t *testing.T) {
	c := service_tester.NewSource(t)
	c.SeedVirtualMachine("alfa")
	s := service_tester.Service(t, c)
	_, e := s.SetVirtualLabel("alfa", "owner", "ops")
	assert.FatalOnError(t, e)
	m, labels, f := s.VirtualMachine("alfa")
	assert.FatalOnError(t, f)
	assert.String(t, "alfa", m.Name)
	assert.Count(t, 1, labels)
	assert.String(t, "ops", labels[0].Value)
}

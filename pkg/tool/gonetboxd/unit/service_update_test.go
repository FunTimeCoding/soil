package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/device_change"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/machine_change"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"testing"
)

func TestUpdateDeviceAppliesOnlyGivenFields(t *testing.T) {
	o := service_tester.New(t)
	c := o.Source
	c.SeedDevice("core")
	v := device_change.New()
	v.Status = new(constant.StatusActive)
	v.Serial = new("S1")
	d, e := o.Service().UpdateDevice("core", v)
	assert.FatalOnError(t, e)
	assert.String(t, "core", d.Name)
	assert.Strings(
		t,
		[]string{"SetDeviceStatus core active", "SetDeviceSerial core S1"},
		c.Calls(),
	)
}

func TestUpdateDeviceRenamesLastAndReturnsRenamed(t *testing.T) {
	o := service_tester.New(t)
	c := o.Source
	c.SeedDevice("core")
	v := device_change.New()
	v.Name = new("edge")
	v.Status = new(constant.StatusOffline)
	d, e := o.Service().UpdateDevice("core", v)
	assert.FatalOnError(t, e)
	assert.String(t, "edge", d.Name)
	assert.Strings(
		t,
		[]string{"SetDeviceStatus core offline", "RenameDevice core edge"},
		c.Calls(),
	)
}

func TestUpdateDeviceRefusesUnknownDevice(t *testing.T) {
	v := device_change.New()
	v.Status = new(constant.StatusActive)
	_, e := service_tester.New(t).Service().UpdateDevice("missing", v)
	assert.True(t, not_found.Is(e))
}

func TestUpdateVirtualMachineSetsRole(t *testing.T) {
	o := service_tester.New(t)
	c := o.Source
	c.SeedVirtualMachine("alfa")
	v := machine_change.New()
	v.Role = new("bravo")
	m, e := o.Service().UpdateVirtualMachine("alfa", v)
	assert.FatalOnError(t, e)
	assert.String(t, "alfa", m.Name)
	assert.Strings(t, []string{"SetVirtualMachineRole alfa bravo"}, c.Calls())
}

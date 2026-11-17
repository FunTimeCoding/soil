package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"testing"
)

func TestCreateDeviceWithTenantAndTags(t *testing.T) {
	o := service_tester.New(t)
	service_tester.SeedDeviceCatalog(o.Source)
	o.Source.SeedTenant("acme")
	d, e := o.Service().CreateDevice(
		"core",
		"server",
		"r640",
		"home",
		"acme",
		[]string{"prod"},
	)
	assert.FatalOnError(t, e)
	assert.String(t, "acme", d.Tenant)
	assert.Strings(t, []string{"prod"}, d.Tags)
}

func TestCreateDeviceWithoutTenant(t *testing.T) {
	o := service_tester.New(t)
	service_tester.SeedDeviceCatalog(o.Source)
	d, e := o.Service().CreateDevice("core", "server", "r640", "home", "", nil)
	assert.FatalOnError(t, e)
	assert.String(t, "no tenant", d.Tenant)
}

func TestCreateDeviceRefusesUnknownRoleBeforeCreating(t *testing.T) {
	o := service_tester.New(t)
	c := o.Source
	c.SeedDeviceType("r640")
	c.SeedSite("home")
	s := o.Service()
	_, e := s.CreateDevice("core", "server", "r640", "home", "", nil)
	assert.True(t, not_found.Is(e))
	assert.String(t, "device role not found: server", e.Error())
	devices, _, f := s.Devices("")
	assert.FatalOnError(t, f)
	assert.Count(t, 0, devices)
}

func TestCreateInterfaceAttachesPhysicalAddress(t *testing.T) {
	o := service_tester.New(t)
	c := o.Source
	c.SeedDevice("core")
	s := o.Service()
	i, e := s.CreateInterface("core", "eno1", "1000base-t", "02:00:00:00:00:01")
	assert.FatalOnError(t, e)
	assert.String(t, "eno1", i.Name)
	assert.Strings(
		t,
		[]string{"CreatePhysicalInterface eno1 02:00:00:00:00:01"},
		c.Calls(),
	)
}

func TestCreateInterfaceWithoutPhysicalAddress(t *testing.T) {
	o := service_tester.New(t)
	c := o.Source
	c.SeedDevice("core")
	_, e := o.Service().CreateInterface("core", "eno1", "1000base-t", "")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, c.Calls())
}

func TestCreateInterfaceRefusesBadPhysicalAddressBeforeCreating(t *testing.T) {
	o := service_tester.New(t)
	c := o.Source
	c.SeedDevice("core")
	s := o.Service()
	_, e := s.CreateInterface("core", "eno1", "1000base-t", "nonsense")
	assert.True(t, validation.Is(e))
	interfaces, f := s.DeviceInterfaces("core")
	assert.FatalOnError(t, f)
	assert.Count(t, 0, interfaces)
}

func TestCreateCableResolvesBothEnds(t *testing.T) {
	o := service_tester.New(t)
	c := o.Source
	c.SeedInterface(c.SeedDevice("core"), "eno1")
	c.SeedInterface(c.SeedDevice("edge"), "eno2")
	created, e := o.Service().CreateCable("core", "eno1", "edge", "eno2")
	assert.FatalOnError(t, e)
	assert.NotNil(t, created)
	assert.Strings(t, []string{"CreateCable eno1 eno2"}, c.Calls())
}

func TestCreateCableRefusesUnknownInterface(t *testing.T) {
	o := service_tester.New(t)
	c := o.Source
	c.SeedInterface(c.SeedDevice("core"), "eno1")
	c.SeedDevice("edge")
	_, e := o.Service().CreateCable("core", "eno1", "edge", "eno2")
	assert.True(t, not_found.Is(e))
	assert.String(t, "interface eno2 not found for device edge", e.Error())
	assert.Count(t, 0, c.Calls())
}

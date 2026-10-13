package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"testing"
)

func TestCreateDeviceWithTenantAndTags(t *testing.T) {
	c := deviceCatalog()
	c.SeedTenant("acme")
	d, e := service_tester.Service(t, c).CreateDevice(
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
	d, e := service_tester.Service(t, deviceCatalog()).CreateDevice(
		"core",
		"server",
		"r640",
		"home",
		"",
		nil,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "no tenant", d.Tenant)
}

func TestCreateDeviceRefusesUnknownRoleBeforeCreating(t *testing.T) {
	c := mock_client.New()
	c.SeedDeviceType("r640")
	c.SeedSite("home")
	s := service_tester.Service(t, c)
	_, e := s.CreateDevice("core", "server", "r640", "home", "", nil)
	assert.True(t, not_found.Is(e))
	assert.String(t, "device role not found: server", e.Error())
	devices, _, f := s.Devices("")
	assert.FatalOnError(t, f)
	assert.Count(t, 0, devices)
}

func TestCreateInterfaceAttachesPhysicalAddress(t *testing.T) {
	c := mock_client.New()
	c.SeedDevice("core")
	s := service_tester.Service(t, c)
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
	c := mock_client.New()
	c.SeedDevice("core")
	_, e := service_tester.Service(t, c).CreateInterface(
		"core",
		"eno1",
		"1000base-t",
		"",
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, c.Calls())
}

func TestCreateInterfaceRefusesBadPhysicalAddressBeforeCreating(t *testing.T) {
	c := mock_client.New()
	c.SeedDevice("core")
	s := service_tester.Service(t, c)
	_, e := s.CreateInterface("core", "eno1", "1000base-t", "nonsense")
	assert.True(t, validation.Is(e))
	interfaces, f := s.DeviceInterfaces("core")
	assert.FatalOnError(t, f)
	assert.Count(t, 0, interfaces)
}

func TestCreateCableResolvesBothEnds(t *testing.T) {
	c := mock_client.New()
	c.SeedInterface(c.SeedDevice("core"), "eno1")
	c.SeedInterface(c.SeedDevice("edge"), "eno2")
	created, e := service_tester.Service(t, c).CreateCable(
		"core",
		"eno1",
		"edge",
		"eno2",
	)
	assert.FatalOnError(t, e)
	assert.NotNil(t, created)
	assert.Strings(t, []string{"CreateCable eno1 eno2"}, c.Calls())
}

func TestCreateCableRefusesUnknownInterface(t *testing.T) {
	c := mock_client.New()
	c.SeedInterface(c.SeedDevice("core"), "eno1")
	c.SeedDevice("edge")
	_, e := service_tester.Service(t, c).CreateCable(
		"core",
		"eno1",
		"edge",
		"eno2",
	)
	assert.True(t, not_found.Is(e))
	assert.String(t, "interface eno2 not found for device edge", e.Error())
	assert.Count(t, 0, c.Calls())
}

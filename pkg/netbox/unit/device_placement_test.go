package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/fixture"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestDeviceTypePatchCarriesTypeIdentifier(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	_, e := c.SetDeviceType("bravo", "november")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.MapValue(t, 9.0, r.Patched, "device_type")
}

func TestDeviceRolePatchCarriesRoleIdentifier(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	_, e := c.SetDeviceRole("bravo", "oscar")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.MapValue(t, 10.0, r.Patched, "role")
}

func TestDeviceRoleRenameCarriesDerivedSlug(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	result, e := c.RenameDeviceRole("oscar", "Quebec Romeo")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.MapValue(t, "Quebec Romeo", r.Patched, constant.NameField)
	assert.MapValue(t, "quebec-romeo", r.Patched, "slug")
	assert.String(t, "Quebec Romeo", result.Name)
}

func TestDeviceSitePatchCarriesSiteIdentifier(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	_, e := c.SetDeviceSite("bravo", "papa")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.MapValue(t, 11.0, r.Patched, constant.SiteField)
}

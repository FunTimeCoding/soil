package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/fixture"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestInterfaceEthernetPowerRead(t *testing.T) {
	s := fixture.NewServer(fixture.NewRecorder())
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	d, e := c.Device("bravo")
	assert.FatalOnError(t, e)
	i, e := c.DeviceInterface(d, "lima")
	assert.FatalOnError(t, e)
	assert.String(t, "pd", i.EthernetPowerMode)
	assert.String(t, "type1-ieee802.3af", i.EthernetPowerType)
}

func TestInterfaceEthernetPowerPatchCarriesModeAndType(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	d, e := c.Device("bravo")
	assert.FatalOnError(t, e)
	_, e = c.SetInterfaceEthernetPower(d, "lima", "pse", "type2-ieee802.3at")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.MapValue(t, "pse", r.Patched, "poe_mode")
	assert.MapValue(t, "type2-ieee802.3at", r.Patched, "poe_type")
}

func TestEmptyInterfaceEthernetPowerSendsNull(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	d, e := c.Device("bravo")
	assert.FatalOnError(t, e)
	_, e = c.SetInterfaceEthernetPower(d, "lima", "", "")
	assert.FatalOnError(t, e)
	assert.MapValue(t, nil, r.Patched, "poe_mode")
	assert.MapValue(t, nil, r.Patched, "poe_type")
}

func TestUnknownEthernetPowerModeRefusedBeforeRequest(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	d, e := c.Device("bravo")
	assert.FatalOnError(t, e)
	_, e = c.SetInterfaceEthernetPower(d, "lima", "mike", "")
	assert.True(t, validation.Is(e))
	assert.Nil(t, r.Patched)
}

func TestDeleteCableSendsDelete(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	assert.FatalOnError(t, c.DeleteCable(7))
	assert.Strings(t, []string{"/api/dcim/cables/7/"}, r.Deleted)
}

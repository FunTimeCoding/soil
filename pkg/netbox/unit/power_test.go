package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/fixture"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestDevicePowerOutletConnectionRead(t *testing.T) {
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
	o, e := c.DevicePowerOutlet(d, "india")
	assert.FatalOnError(t, e)
	assert.String(t, "bravo", o.Device)
	assert.String(t, "kilo", o.ConnectedDevice)
	assert.String(t, "juliett", o.ConnectedPort)
}

func TestCreatePowerOutletCarriesDeviceAndName(t *testing.T) {
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
	_, e = c.CreatePowerOutlet(d, "india")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.MapValue(t, "india", r.Created, constant.NameField)
	assert.MapValue(
		t,
		map[string]any{"name": "bravo"},
		r.Created,
		constant.DeviceField,
	)
}

func TestCreatePowerPortCarriesDeviceAndName(t *testing.T) {
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
	p, e := c.CreatePowerPort(d, "juliett")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.MapValue(t, "juliett", r.Created, constant.NameField)
	assert.MapValue(
		t,
		map[string]any{"name": "bravo"},
		r.Created,
		constant.DeviceField,
	)
	assert.String(t, "kilo", p.Device)
}

func TestPowerCableJoinsOutletToPort(t *testing.T) {
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
	o, e := c.DevicePowerOutlet(d, "india")
	assert.FatalOnError(t, e)
	p, e := c.CreatePowerPort(d, "juliett")
	assert.FatalOnError(t, e)
	_, e = c.CreatePowerCable(o, p)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.MapValue(
		t,
		[]any{
			map[string]any{"object_type": "dcim.poweroutlet", "object_id": 5.0},
		},
		r.Created,
		"a_terminations",
	)
	assert.MapValue(
		t,
		[]any{map[string]any{"object_type": "dcim.powerport", "object_id": 6.0}},
		r.Created,
		"b_terminations",
	)
}

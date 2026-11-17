package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/fixture"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestDeviceLocalContextRead(t *testing.T) {
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
	assert.Any(t, map[string]any{"golf": "hotel"}, d.LocalContext)
}

func TestDeviceLocalContextPatchCarriesPayload(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	_, e := c.SetDeviceLocalContext("bravo", map[string]any{"india": "juliett"})
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.MapValue(
		t,
		map[string]any{"india": "juliett"},
		r.Patched,
		"local_context_data",
	)
}

func TestEmptyDeviceLocalContextPatchSendsNull(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	_, e := c.SetDeviceLocalContext("bravo", map[string]any{})
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.MapValue(t, nil, r.Patched, "local_context_data")
}

func TestVirtualMachineLocalContextRead(t *testing.T) {
	s := fixture.NewServer(fixture.NewRecorder())
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	m, e := c.VirtualMachine("alfa")
	assert.FatalOnError(t, e)
	assert.Any(t, map[string]any{"golf": "hotel"}, m.LocalContext)
}

func TestVirtualMachineLocalContextPatchCarriesPayload(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	_, e := c.SetVirtualMachineLocalContext(
		"alfa",
		map[string]any{"india": "juliett"},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.MapValue(
		t,
		map[string]any{"india": "juliett"},
		r.Patched,
		"local_context_data",
	)
}

func TestEmptyVirtualMachineLocalContextPatchSendsNull(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	_, e := c.SetVirtualMachineLocalContext("alfa", map[string]any{})
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.MapValue(t, nil, r.Patched, "local_context_data")
}

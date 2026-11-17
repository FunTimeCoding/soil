package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"net/http"
	"testing"
)

func TestDeviceLocalContextRoundTrip(t *testing.T) {
	x := context.Background()
	s := base.New(t)
	s.Source.SeedDevice("bravo")
	c := s.GeneratedClient()
	set, e := c.UpdateDeviceWithResponse(
		x,
		"bravo",
		client.UpdateDeviceRequest{
			LocalContext: &map[string]any{"charlie": "delta"},
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, set.StatusCode())
	assert.Any(t, &map[string]any{"charlie": "delta"}, set.JSON200.LocalContext)
	cleared, e := c.UpdateDeviceWithResponse(
		x,
		"bravo",
		client.UpdateDeviceRequest{LocalContext: &map[string]any{}},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, cleared.StatusCode())
	assert.Nil(t, cleared.JSON200.LocalContext)
}

func TestDeviceLocalContextUnchangedByOtherUpdate(t *testing.T) {
	x := context.Background()
	s := base.New(t)
	s.Source.SeedDevice("bravo")
	c := s.GeneratedClient()
	_, e := c.UpdateDeviceWithResponse(
		x,
		"bravo",
		client.UpdateDeviceRequest{
			LocalContext: &map[string]any{"charlie": "delta"},
		},
	)
	assert.FatalOnError(t, e)
	described, e := c.UpdateDeviceWithResponse(
		x,
		"bravo",
		client.UpdateDeviceRequest{Description: new("echo")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, described.StatusCode())
	assert.Any(
		t,
		&map[string]any{"charlie": "delta"},
		described.JSON200.LocalContext,
	)
}

func TestVirtualMachineLocalContextRoundTrip(t *testing.T) {
	x := context.Background()
	s := base.New(t)
	s.Source.SeedVirtualMachine("charlie")
	c := s.GeneratedClient()
	set, e := c.UpdateVirtualMachineWithResponse(
		x,
		"charlie",
		client.UpdateVirtualMachineRequest{
			LocalContext: &map[string]any{"delta": "echo"},
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, set.StatusCode())
	assert.Any(t, &map[string]any{"delta": "echo"}, set.JSON200.LocalContext)
	cleared, e := c.UpdateVirtualMachineWithResponse(
		x,
		"charlie",
		client.UpdateVirtualMachineRequest{LocalContext: &map[string]any{}},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, cleared.StatusCode())
	assert.Nil(t, cleared.JSON200.LocalContext)
}

package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"net/http"
	"testing"
)

func TestPowerCableShowsOnOutletList(t *testing.T) {
	x := context.Background()
	s := base.New(t)
	s.Source.SeedDevice("alfa")
	s.Source.SeedDevice("bravo")
	c := s.GeneratedClient()
	outlet, e := c.CreatePowerOutletWithResponse(
		x,
		"alfa",
		client.CreatePowerOutletRequest{Name: "charlie"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusCreated, outlet.StatusCode())
	port, e := c.CreatePowerPortWithResponse(
		x,
		"bravo",
		client.CreatePowerPortRequest{Name: "delta"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusCreated, port.StatusCode())
	cable, e := c.CreatePowerCableWithResponse(
		x,
		client.CreatePowerCableRequest{
			OutletDevice: "alfa",
			Outlet:       "charlie",
			PortDevice:   "bravo",
			Port:         "delta",
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusCreated, cable.StatusCode())
	listed, e := c.ListPowerOutletsWithResponse(x, "alfa")
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, listed.StatusCode())
	assert.Count(t, 1, *listed.JSON200)
	o := (*listed.JSON200)[0]
	assert.String(t, "charlie", o.Name)
	assert.String(t, "alfa", o.Device)
	assert.Any(t, new("bravo"), o.ConnectedDevice)
	assert.Any(t, new("delta"), o.ConnectedPort)
}

func TestUnpluggedOutletListedWithoutConnection(t *testing.T) {
	x := context.Background()
	s := base.New(t)
	s.Source.SeedDevice("alfa")
	c := s.GeneratedClient()
	_, e := c.CreatePowerOutletWithResponse(
		x,
		"alfa",
		client.CreatePowerOutletRequest{Name: "charlie"},
	)
	assert.FatalOnError(t, e)
	listed, e := c.ListPowerOutletsWithResponse(x, "alfa")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *listed.JSON200)
	assert.Nil(t, (*listed.JSON200)[0].ConnectedDevice)
}

func TestPowerCableToMissingPortNotFound(t *testing.T) {
	x := context.Background()
	s := base.New(t)
	s.Source.SeedDevice("alfa")
	s.Source.SeedDevice("bravo")
	c := s.GeneratedClient()
	_, e := c.CreatePowerOutletWithResponse(
		x,
		"alfa",
		client.CreatePowerOutletRequest{Name: "charlie"},
	)
	assert.FatalOnError(t, e)
	cable, e := c.CreatePowerCableWithResponse(
		x,
		client.CreatePowerCableRequest{
			OutletDevice: "alfa",
			Outlet:       "charlie",
			PortDevice:   "bravo",
			Port:         "delta",
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, cable.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

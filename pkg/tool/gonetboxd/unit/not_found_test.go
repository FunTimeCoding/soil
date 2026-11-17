package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"net/http"
	"testing"
)

func TestListedRoleCarriesVirtualMachineFlag(t *testing.T) {
	s := base.New(t)
	s.Source.SeedVirtualMachineRole("golf")
	response, e := s.GeneratedClient().ListDeviceRolesWithResponse(
		context.Background(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *response.JSON200)
	assert.True(t, (*response.JSON200)[0].VirtualMachine)
}

func TestMissingDeviceIsNotFoundWithoutSentry(t *testing.T) {
	s := base.New(t)
	response, e := s.GeneratedClient().GetDeviceWithResponse(
		context.Background(),
		"alfa",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestMissingVirtualMachineIsNotFoundWithoutSentry(t *testing.T) {
	s := base.New(t)
	response, e := s.GeneratedClient().GetVirtualMachineWithResponse(
		context.Background(),
		"alfa",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestUpdateMissingVirtualMachineIsNotFound(t *testing.T) {
	s := base.New(t)
	response, e := s.GeneratedClient().UpdateVirtualMachineWithResponse(
		context.Background(),
		"alfa",
		client.UpdateVirtualMachineRequest{Role: new("bravo")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestInterfacesOfMissingDeviceAreNotFound(t *testing.T) {
	s := base.New(t)
	response, e := s.GeneratedClient().ListInterfacesWithResponse(
		context.Background(),
		"alfa",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestTunnelTerminationOnMissingDeviceIsNotFound(t *testing.T) {
	s := base.New(t)
	response, e := s.GeneratedClient().CreateDeviceTunnelTerminationWithResponse(
		context.Background(),
		"alfa",
		client.CreateTunnelTerminationRequest{
			Interface: "bravo",
			Role:      "peer",
			Tunnel:    "charlie",
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestLabelOfMissingDeviceIsNotFound(t *testing.T) {
	s := base.New(t)
	response, e := s.GeneratedClient().RemoveDeviceLabelWithResponse(
		context.Background(),
		"alfa",
		"bravo",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestVirtualMachineTagOnMissingMachineIsNotFound(t *testing.T) {
	s := base.New(t)
	response, e := s.GeneratedClient().AddVirtualTagWithResponse(
		context.Background(),
		"alfa",
		"bravo",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

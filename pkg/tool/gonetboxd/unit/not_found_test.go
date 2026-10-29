package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"net/http"
	"testing"
)

func TestListedRoleCarriesVirtualMachineFlag(t *testing.T) {
	m := service_tester.NewSource(t)
	m.SeedVirtualMachineRole("golf")
	c := newClient(t, newNetboxServer(t, m, memory.New()))
	response, e := c.ListDeviceRolesWithResponse(context.Background())
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *response.JSON200)
	assert.True(t, (*response.JSON200)[0].VirtualMachine)
}

func TestMissingDeviceIsNotFoundWithoutSentry(t *testing.T) {
	r := memory.New()
	c := newClient(t, newNetboxServer(t, service_tester.NewSource(t), r))
	response, e := c.GetDeviceWithResponse(context.Background(), "alfa")
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, r.Events())
}

func TestMissingVirtualMachineIsNotFoundWithoutSentry(t *testing.T) {
	r := memory.New()
	c := newClient(t, newNetboxServer(t, service_tester.NewSource(t), r))
	response, e := c.GetVirtualMachineWithResponse(context.Background(), "alfa")
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, r.Events())
}

func TestUpdateMissingVirtualMachineIsNotFound(t *testing.T) {
	r := memory.New()
	c := newClient(t, newNetboxServer(t, service_tester.NewSource(t), r))
	response, e := c.UpdateVirtualMachineWithResponse(
		context.Background(),
		"alfa",
		client.UpdateVirtualMachineRequest{Role: new("bravo")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, r.Events())
}

func TestInterfacesOfMissingDeviceAreNotFound(t *testing.T) {
	r := memory.New()
	c := newClient(t, newNetboxServer(t, service_tester.NewSource(t), r))
	response, e := c.ListInterfacesWithResponse(context.Background(), "alfa")
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, r.Events())
}

func TestTunnelTerminationOnMissingDeviceIsNotFound(t *testing.T) {
	r := memory.New()
	c := newClient(t, newNetboxServer(t, service_tester.NewSource(t), r))
	response, e := c.CreateDeviceTunnelTerminationWithResponse(
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
	assert.Count(t, 0, r.Events())
}

func TestLabelOfMissingDeviceIsNotFound(t *testing.T) {
	r := memory.New()
	c := newClient(t, newNetboxServer(t, service_tester.NewSource(t), r))
	response, e := c.RemoveDeviceLabelWithResponse(
		context.Background(),
		"alfa",
		"bravo",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, r.Events())
}

func TestVirtualMachineTagOnMissingMachineIsNotFound(t *testing.T) {
	r := memory.New()
	c := newClient(t, newNetboxServer(t, service_tester.NewSource(t), r))
	response, e := c.AddVirtualTagWithResponse(
		context.Background(),
		"alfa",
		"bravo",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, r.Events())
}

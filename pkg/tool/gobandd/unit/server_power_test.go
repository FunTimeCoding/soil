package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func TestARefusedPowerRequestSendsNothingToTheEngine(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(bandConstant.PowerOn, hardStates()...)
	r, e := s.Client.RequestPowerWithResponse(
		t.Context(),
		client.RequestPowerJSONRequestBody{
			Instance:  new("alfa"),
			Operation: client.PowerRequestOperationShutdown,
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 409, r.StatusCode())
	assert.StringContains(t, "off soft graceful not offered", r.JSON409.Error)
	assert.Count(t, 0, s.Alfa.Requested())
}

func TestAGracefulShutdownRequestsOffSoftGraceful(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(bandConstant.PowerOn, runningStates()...)
	r, e := s.Client.RequestPowerWithResponse(
		t.Context(),
		client.RequestPowerJSONRequestBody{
			Instance:  new("alfa"),
			Operation: client.PowerRequestOperationShutdown,
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.String(t, "off soft graceful", r.JSON200.Requested)
	assert.String(t, "on", r.JSON200.Power)
	assert.Integers(t, []int{12}, s.Alfa.Requested())
}

func TestAForcedRebootRequestsAPowerCycle(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(bandConstant.PowerOn, runningStates()...)
	r, e := s.Client.RequestPowerWithResponse(
		t.Context(),
		client.RequestPowerJSONRequestBody{
			Instance:  new("alfa"),
			Operation: client.PowerRequestOperationReboot,
			Force:     new(true),
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.Integers(t, []int{5}, s.Alfa.Requested())
	assert.Count(t, 0, s.Bravo.Requested())
}

func TestPowerWithoutAnInstanceIsRefusedWhileTwoAreConfigured(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetPower(bandConstant.PowerOn, runningStates()...)
	s.Bravo.SetPower(bandConstant.PowerOn, runningStates()...)
	r, e := s.Client.RequestPowerWithResponse(
		t.Context(),
		client.RequestPowerJSONRequestBody{
			Operation: client.PowerRequestOperationReboot,
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 400, r.StatusCode())
	assert.Count(t, 0, s.Alfa.Requested())
	assert.Count(t, 0, s.Bravo.Requested())
}

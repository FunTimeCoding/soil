package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestCruiseRoundTripsThroughRest(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	r, e := a.RestClient.PostSessionCruiseWithResponse(
		a.Context,
		a.UUID,
		client.CruiseRequest{Mode: client.Paced, PaceMinutes: new(15)},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	g, f := a.RestClient.GetSessionCruiseWithResponse(a.Context, a.UUID)
	assert.FatalOnError(t, f)
	assert.String(t, "paced", g.JSON200.Mode)
	assert.Integer(t, 15, *g.JSON200.PaceMinutes)
	assert.Integer(t, 0, g.JSON200.Woken)
}

func TestRestRefusesPacedCruiseWithoutAPace(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	r, e := a.RestClient.PostSessionCruiseWithResponse(
		a.Context,
		a.UUID,
		client.CruiseRequest{Mode: client.Paced},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 400, r.StatusCode())
}

func TestCruisePulseOverRestReportsTheHold(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.RestClient.PostSessionPulseWithResponse(
		a.Context,
		a.UUID,
		client.PulseRequest{Body: "alert batch", Delivery: new(client.Cruise)},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.False(t, r.JSON200.Immediate)
	pulses := clientEntriesByKind(a.CheckLive().Entries, constant.QueuePulse)
	assert.Count(t, 1, pulses)
}

func TestCruiseToolSetsAndReadsTheOwnSession(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.CruiseTool,
		map[string]any{constant.Mode: "paced", constant.Pace: 30},
	)
	assert.String(
		t,
		"cruise paced every 30m0s",
		a.MustCallTool(constant.CruiseTool, map[string]any{}),
	)
}

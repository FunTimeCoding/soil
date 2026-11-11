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
		client.CruiseRequest{Grace: "15"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	g, f := a.RestClient.GetSessionCruiseWithResponse(a.Context, a.UUID)
	assert.FatalOnError(t, f)
	assert.Integer(t, 15, g.JSON200.GraceMinutes)
	assert.False(t, g.JSON200.Default)
	assert.Integer(t, 0, g.JSON200.Woken)
}

func TestRestReportsDefaultCruise(t *testing.T) {
	s := base.NewWithChannelAlways(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	g, e := a.RestClient.GetSessionCruiseWithResponse(a.Context, a.UUID)
	assert.FatalOnError(t, e)
	assert.Integer(t, 30, g.JSON200.GraceMinutes)
	assert.True(t, g.JSON200.Default)
}

func TestRestRefusesGracePastLimit(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	r, e := a.RestClient.PostSessionCruiseWithResponse(
		a.Context,
		a.UUID,
		client.CruiseRequest{Grace: "60"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 400, r.StatusCode())
}

func TestCruiseToolSetsAndReadsOwnSession(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.MustCallTool(constant.CruiseTool, map[string]any{constant.Grace: "20"})
	assert.String(
		t,
		"cruise after 20m0s idle",
		a.MustCallTool(constant.CruiseTool, map[string]any{}),
	)
}

func TestCruiseToolRefusesGracePastLimit(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	answer := a.MustCallToolError(
		constant.CruiseTool,
		map[string]any{constant.Grace: "61"},
	)
	assert.StringContains(t, "from 1 to 59", answer)
}

package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
	"time"
)

func TestPacedCruiseTakesOneToSixtyMinutes(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	assert.FatalOnError(
		t,
		s.Message.SetCruise("session-1", constant.CruisePaced, time.Minute),
	)
	assert.FatalOnError(
		t,
		s.Message.SetCruise("session-1", constant.CruisePaced, time.Hour),
	)
}

func TestPacedCruiseRefusesPastHour(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	e := s.Message.SetCruise("session-1", constant.CruisePaced, 61*time.Minute)
	assert.True(t, refusal.Is(e))
	assert.StringContains(t, "from 1 to 60, not 61", e.Error())
}

func TestPacedCruiseRefusesPartOfMinute(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	e := s.Message.SetCruise("session-1", constant.CruisePaced, 90*time.Second)
	assert.True(t, refusal.Is(e))
	assert.StringContains(t, "whole number of minutes", e.Error())
}

func TestCruiseToolRefusesFractionalPace(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	result := a.MustCallToolError(
		constant.CruiseTool,
		map[string]any{constant.Mode: "paced", constant.Pace: 2.5},
	)
	assert.StringContains(t, "from 1 to 60, not 2.5", result)
}

func TestRestRefusesPacePastHour(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	r, e := a.RestClient.PostSessionCruiseWithResponse(
		a.Context,
		a.UUID,
		client.CruiseRequest{Mode: client.Paced, PaceMinutes: new(61)},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 400, r.StatusCode())
}

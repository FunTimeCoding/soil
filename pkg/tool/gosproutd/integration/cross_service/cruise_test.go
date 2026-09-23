package cross_service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration/cross_service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/pulse"
	"testing"
	"time"
)

func TestArmedCruiseWakesAnIdleSession(t *testing.T) {
	s := cross_service_tester.New(t)
	a := s.Goclauded.NewSession(t)
	defer a.Close()
	a.CheckLive()
	s.Goclauded.Store.Advance(time.Minute)
	s.EndTurn(a.UUID)
	s.AnswerOne(stringConstant.LowerAlfa)
	s.Service.SetCruise(stringConstant.LowerAlfa, constant.CruiseOn, 0)
	s.Advance(constant.QuietWindow + time.Minute)
	assert.String(
		t,
		"immediate",
		string(
			pulse.Decide(
				s.Service.Reading(stringConstant.LowerAlfa, s.Target(a.UUID)),
				s.Now(),
			),
		),
	)
}

func TestDisarmedCruiseLeavesTheSameStateQueued(t *testing.T) {
	s := cross_service_tester.New(t)
	a := s.Goclauded.NewSession(t)
	defer a.Close()
	a.CheckLive()
	s.Goclauded.Store.Advance(time.Minute)
	s.EndTurn(a.UUID)
	s.AnswerOne(stringConstant.LowerAlfa)
	s.Advance(constant.QuietWindow + time.Minute)
	assert.String(
		t,
		"queue",
		string(
			pulse.Decide(
				s.Service.Reading(stringConstant.LowerAlfa, s.Target(a.UUID)),
				s.Now(),
			),
		),
	)
}

func TestArmedCruiseLeavesAWorkingSessionAlone(t *testing.T) {
	s := cross_service_tester.New(t)
	a := s.Goclauded.NewSession(t)
	defer a.Close()
	s.EndTurn(a.UUID)
	s.Goclauded.Store.Advance(time.Minute)
	a.CheckLive()
	s.AnswerOne(stringConstant.LowerAlfa)
	s.Service.SetCruise(stringConstant.LowerAlfa, constant.CruiseOn, 0)
	s.Advance(constant.QuietWindow + time.Minute)
	assert.String(
		t,
		"queue",
		string(
			pulse.Decide(
				s.Service.Reading(stringConstant.LowerAlfa, s.Target(a.UUID)),
				s.Now(),
			),
		),
	)
}

func TestPollingClearsTheBacklogSoNothingFires(t *testing.T) {
	s := cross_service_tester.New(t)
	a := s.Goclauded.NewSession(t)
	defer a.Close()
	a.CheckLive()
	s.Goclauded.Store.Advance(time.Minute)
	s.EndTurn(a.UUID)
	s.AnswerOne(stringConstant.LowerAlfa)
	s.Service.SetCruise(stringConstant.LowerAlfa, constant.CruiseOn, 0)
	s.Service.MarkSeen(stringConstant.LowerAlfa)
	s.Advance(constant.QuietWindow + time.Minute)
	assert.String(
		t,
		"drop",
		string(
			pulse.Decide(
				s.Service.Reading(stringConstant.LowerAlfa, s.Target(a.UUID)),
				s.Now(),
			),
		),
	)
}

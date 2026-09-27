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

func TestCoordinatorReportsWorkingAfterAPrompt(t *testing.T) {
	s := cross_service_tester.New(t)
	a := s.Goclauded.NewSession(t)
	defer a.Close()
	a.CheckLive()
	target := s.Target(a.UUID)
	assert.True(t, target.LastPromptAt != nil)
	assert.True(t, target.LastTurnEndAt == nil)
}

func TestCoordinatorReportsIdleAfterTurnEnd(t *testing.T) {
	s := cross_service_tester.New(t)
	a := s.Goclauded.NewSession(t)
	defer a.Close()
	a.CheckLive()
	s.Goclauded.Store.Advance(time.Minute)
	s.EndTurn(a.UUID)
	target := s.Target(a.UUID)
	assert.True(t, target.LastTurnEndAt != nil)
	assert.True(t, target.LastTurnEndAt.After(*target.LastPromptAt))
}

func TestSessionWithoutTheHookReportsUnknown(t *testing.T) {
	s := cross_service_tester.New(t)
	a := s.Goclauded.NewSession(t)
	defer a.Close()
	a.CheckLive()
	assert.True(t, s.Target(a.UUID).LastTurnEndAt == nil)
}

func TestSproutReachesTheCoordinator(t *testing.T) {
	s := cross_service_tester.New(t)
	a := s.Goclauded.NewSession(t)
	defer a.Close()
	targets, e := s.Coordinator.RecentSessions(10)
	assert.FatalOnError(t, e)
	assert.True(t, len(targets) > 0)
}

func TestCoordinatorSeesOnlyItsOwnSessions(t *testing.T) {
	s := cross_service_tester.New(t)
	before, e := s.Coordinator.RecentSessions(10)
	assert.FatalOnError(t, e)
	a := s.Goclauded.NewSession(t)
	defer a.Close()
	after, f := s.Coordinator.RecentSessions(10)
	assert.FatalOnError(t, f)
	assert.Integer(t, len(before)+1, len(after))
}

func TestQueueAndCoordinatorKeepSeparateClocks(t *testing.T) {
	s := cross_service_tester.New(t)
	a := s.Goclauded.NewSession(t)
	defer a.Close()
	d, e := s.Service.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	assert.Nil(t, e)
	s.Advance(time.Hour)
	s.Service.Answer(d.Identifier, constant.AnswerKindChoice, "narrow")
	answered := s.Service.Decisions(stringConstant.LowerAlfa)[0].AnsweredAt
	assert.True(t, answered != nil)
	assert.True(
		t,
		s.Goclauded.Store.GetSession(a.UUID).LastSeen.Before(*answered),
	)
}

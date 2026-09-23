package cross_service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration/cross_service_tester"
	"testing"
	"time"
)

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

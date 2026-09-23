package cross_service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration/cross_service_tester"
	"testing"
	"time"
)

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

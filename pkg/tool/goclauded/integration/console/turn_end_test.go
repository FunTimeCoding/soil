package console

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/console_tester"
	"testing"
	"time"
)

func TestTurnEndThroughTheHookCommand(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Port)
	c.Register("session-1")
	assert.True(t, s.Store.GetSession("session-1").LastTurnEndAt == nil)
	c.TurnEnd("session-1")
	assert.True(t, s.Store.GetSession("session-1").LastTurnEndAt != nil)
}

func TestHookCommandOrderingDecidesIdle(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Port)
	c.Register("session-1")
	c.Check("session-1")
	s.Store.Advance(time.Minute)
	c.TurnEnd("session-1")
	r := s.Store.GetSession("session-1")
	assert.True(t, r.LastTurnEndAt.After(*r.LastPromptAt))
}

func TestHookCommandOrderingDecidesWorking(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Port)
	c.Register("session-1")
	c.TurnEnd("session-1")
	s.Store.Advance(time.Minute)
	c.Check("session-1")
	r := s.Store.GetSession("session-1")
	assert.True(t, r.LastPromptAt.After(*r.LastTurnEndAt))
}

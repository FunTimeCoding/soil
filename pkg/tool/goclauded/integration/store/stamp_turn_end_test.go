package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
	"time"
)

func TestStampTurnEndRecordsTheBoundary(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	assert.True(t, s.GetSession("session-1").LastTurnEndAt == nil)
	assert.FatalOnError(t, s.Store.StampTurnEnd("session-1", time.Now()))
	assert.True(t, s.GetSession("session-1").LastTurnEndAt != nil)
}

func TestTurnEndBeforePromptMeansWorking(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	base := time.Now()
	assert.FatalOnError(t, s.Store.StampTurnEnd("session-1", base))
	assert.FatalOnError(
		t,
		s.Store.StampPrompt("session-1", base.Add(time.Minute)),
	)
	r := s.GetSession("session-1")
	assert.True(t, r.LastPromptAt.After(*r.LastTurnEndAt))
}

func TestTurnEndAfterPromptMeansIdle(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	base := time.Now()
	assert.FatalOnError(t, s.Store.StampPrompt("session-1", base))
	assert.FatalOnError(
		t,
		s.Store.StampTurnEnd("session-1", base.Add(time.Minute)),
	)
	r := s.GetSession("session-1")
	assert.True(t, r.LastTurnEndAt.After(*r.LastPromptAt))
}

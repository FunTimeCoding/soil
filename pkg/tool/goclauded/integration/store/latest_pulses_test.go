package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
)

func TestLatestPulsesBySessions(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	s.EnsureSession("session-2")
	s.EnsureSession("session-3")
	assert.FatalOnError(t, s.Store.SendPulse("session-1", "Ash", "first"))
	assert.FatalOnError(t, s.Store.SendPulse("session-1", "Ash", "second"))
	assert.FatalOnError(t, s.Store.SendPulse("session-2", "Blair", "only"))
	result, e := s.Store.LatestPulsesBySessions(
		[]string{"session-1", "session-2", "session-3"},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, result)
	assert.String(t, "second", result["session-1"].Body)
	assert.String(t, "only", result["session-2"].Body)
	assert.Nil(t, result["session-3"])
}

func TestLatestPulsesBySessionsAgreesWithFindLatestPulse(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	assert.FatalOnError(t, s.Store.SendPulse("session-1", "Ash", "first"))
	assert.FatalOnError(t, s.Store.SendPulse("session-1", "Ash", "second"))
	single, found, e := s.Store.FindLatestPulse("session-1")
	assert.FatalOnError(t, e)
	assert.True(t, found)
	batch, f := s.Store.LatestPulsesBySessions([]string{"session-1"})
	assert.FatalOnError(t, f)
	assert.String(t, single.Body, batch["session-1"].Body)
}

func TestLatestPulsesBySessionsWithoutIdentifiers(t *testing.T) {
	s := store_tester.New(t)
	result, e := s.Store.LatestPulsesBySessions(nil)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, result)
}

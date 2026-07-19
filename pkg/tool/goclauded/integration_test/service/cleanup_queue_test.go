package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
	"time"
)

func TestCleanupQueueRemovesConsumedEntries(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Announce("session-1", r.Callsign, "working", "")
	s.Check("session-1")
	s.Store.Advance(25 * time.Hour)
	s.Service.RunTimeoutSweep()
	r = s.Check("session-1")
	announces := entriesByKind(r.Entries, constant.QueueSessionAnnounce)
	assert.Count(t, 0, announces)
}

func TestCleanupQueuePreservesUnconsumedEntries(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Announce("session-1", r.Callsign, "working", "")
	s.Store.Advance(25 * time.Hour)
	s.Service.RunTimeoutSweep()
	r = s.Check("session-1")
	announces := entriesByKind(r.Entries, constant.QueueSessionAnnounce)
	assert.True(t, len(announces) > 0)
}

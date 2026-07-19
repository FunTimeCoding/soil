package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestRelease(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.Announce(r.Callsign, "working", "")
	sessions := s.Store.ListSessions()
	assert.Count(t, 1, sessions)
	s.Release("session-1", r.Callsign)
	sessions = s.Store.ListSessions()
	assert.Count(t, 0, sessions)
	events := s.Store.Events(event_query.New().Kind(constant.Release).SetLimit(10))
	assert.Count(t, 1, events)
}

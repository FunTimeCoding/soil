package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
	"time"
)

func TestInactivitySweepSkipsActiveLog(t *testing.T) {
	s := service_tester.New(t)
	r := s.Register("session-1")
	assert.True(t, r.Callsign != "")
	s.Update("session-1", r.Callsign, "working", "", "")
	s.Store.Advance(2 * time.Hour)
	s.Store.Store.UpdateFields(
		"session-1",
		map[string]any{"last_active_at": s.Store.Clock()()},
	)
	s.Service.RunTimeoutSweep()
	e := s.Store.GetSession("session-1")
	assert.String(t, "", e.TimedOut)
}

func TestInactivitySweepFiresOnStaleLog(t *testing.T) {
	s := service_tester.New(t)
	r := s.Register("session-1")
	assert.True(t, r.Callsign != "")
	s.Update("session-1", r.Callsign, "working", "", "")
	s.Store.Store.UpdateFields(
		"session-1",
		map[string]any{"last_active_at": s.Store.Clock()()},
	)
	s.Store.Advance(2 * time.Hour)
	s.Service.RunTimeoutSweep()
	e := s.Store.GetSession("session-1")
	assert.True(t, e.TimedOut != "")
}

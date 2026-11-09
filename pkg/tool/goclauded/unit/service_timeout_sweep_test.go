package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
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
	assert.Count(t, 0, s.Roster.RunTimeoutSweep())
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
	assert.Count(t, 0, s.Roster.RunTimeoutSweep())
	e := s.Store.GetSession("session-1")
	assert.True(t, e.TimedOut != "")
}

func TestHalfHourQuietIsNoTimeout(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Announce("session-1", r.Callsign, "still working", "")
	s.Check("session-1")
	s.Store.Advance(30 * time.Minute)
	assert.Count(t, 0, s.Roster.RunTimeoutSweep())
	r = s.Check("session-1")
	assert.Count(t, 0, fixture.KindEntries(r.Entries, constant.QueueTimeout))
}

func TestTimeoutIsDeliveredOnce(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Announce("session-1", r.Callsign, "working", "")
	s.Check("session-1")
	s.Store.Advance(2 * time.Hour)
	assert.Count(t, 0, s.Roster.RunTimeoutSweep())
	first := s.Check("session-1")
	assert.Count(
		t,
		1,
		fixture.KindEntries(first.Entries, constant.QueueTimeout),
	)
	second := s.Check("session-1")
	assert.Count(
		t,
		0,
		fixture.KindEntries(second.Entries, constant.QueueTimeout),
	)
}

func TestTimeoutIsLoggedForOthers(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Announce("session-1", r.Callsign, "going quiet", "")
	s.Store.Advance(2 * time.Hour)
	assert.Count(t, 0, s.Roster.RunTimeoutSweep())
	events := s.Store.Events(
		event_query.New().Kind(constant.InactivityTimeout).SetLimit(10),
	)
	assert.Count(t, 1, events)
	assert.String(t, r.Callsign, events[0].Actor)
}

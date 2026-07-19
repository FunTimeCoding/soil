package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"testing"
	"time"
)

func entriesByKind(
	entries []queue.Entry,
	kind string,
) []queue.Entry {
	var result []queue.Entry

	for _, e := range entries {
		if e.Kind == kind {
			result = append(result, e)
		}
	}

	return result
}

func TestCheckNewSession(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	assert.True(t, r.Callsign != "")
}

func TestCheckNoChanges(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	second := s.Check("session-1")
	assert.False(t, second.Changed)
}

func TestCheckWithMessage(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	s.Store.EnsureSession("session-2")
	s.Store.Advance(time.Second)
	s.Send(r1.Callsign, "", "broadcast")
	r := s.Check("session-1")
	assert.True(t, r.Changed)
	messages := entriesByKind(r.Entries, constant.QueueMessage)
	assert.Count(t, 1, messages)
}

func TestCheckWithTimeout(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Announce("session-1", r.Callsign, "working", "")
	s.Check("session-1")
	s.Store.Advance(2 * time.Hour)
	s.Service.RunTimeoutSweep()
	r = s.Check("session-1")
	assert.True(t, r.Changed)
	timeouts := entriesByKind(r.Entries, constant.QueueTimeout)
	assert.Count(t, 1, timeouts)
	assert.String(
		t,
		"1 hour since last turn. Removed from roster. Last topic: working",
		timeouts[0].Body,
	)
}

func TestCheckWithReannounce(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Store.BindModelContextSession(r.Callsign, "mcp-abc")
	s.Service.ClearBindings()
	r = s.Check("session-1")
	assert.True(t, r.Changed)
	reannounce := entriesByKind(r.Entries, constant.QueueReannounce)
	assert.Count(t, 1, reannounce)
}

func TestCheckWithSessionActivity(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Store.Advance(time.Second)
	s.Announce("session-2", r2.Callsign, "other work", "pkg/auth")
	r := s.Check("session-1")
	assert.True(t, r.Changed)
	announces := entriesByKind(r.Entries, constant.QueueSessionAnnounce)
	assert.True(t, len(announces) > 0)
}

func TestCheckCompleteTimeout(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Announce("session-1", r.Callsign, constant.FixtureTopic, "")
	s.Check("session-1")
	s.Store.CompleteTask(r.Callsign)
	s.Store.Advance(time.Hour)
	s.Service.RunTimeoutSweep()
	r = s.Check("session-1")
	assert.True(t, r.Changed)
	timeouts := entriesByKind(r.Entries, constant.QueueTimeout)
	assert.Count(t, 1, timeouts)
	assert.String(
		t,
		"30 minutes since completing. Removed from roster.",
		timeouts[0].Body,
	)
}

func TestCheckConsumesEntries(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Send(r2.Callsign, r1.Callsign, "hello")
	first := s.Check("session-1")
	messages := entriesByKind(first.Entries, constant.QueueMessage)
	assert.Count(t, 1, messages)
	second := s.Check("session-1")
	messages = entriesByKind(second.Entries, constant.QueueMessage)
	assert.Count(t, 0, messages)
}

func TestCheckWithCompletionActivity(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Announce("session-2", r2.Callsign, "some work", "")
	s.Check("session-1")
	s.Complete("session-2", r2.Callsign, "some work", "done")
	s.Store.Advance(time.Second)
	r := s.Check("session-1")
	assert.True(t, r.Changed)
	completions := entriesByKind(r.Entries, constant.QueueSessionComplete)
	assert.True(t, len(completions) > 0)
}

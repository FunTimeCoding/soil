package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
	"time"
)

func TestInactivityTimeout(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working on something")
	a.CheckLive()
	s.Store.Advance(2 * time.Hour)
	s.Service.RunTimeoutSweep()
	r := a.CheckLive()
	timeouts := clientEntriesByKind(r.Entries, constant.QueueTimeout)
	assert.Count(t, 1, timeouts)
	assert.String(
		t,
		"1 hour since last turn. Removed from roster. Last topic: working on something",
		timeouts[0].Body,
	)
}

func TestCompleteTimeoutCoordination(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "a task")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "done",
		},
	)
	a.CheckLive()
	s.Store.Advance(45 * time.Minute)
	s.Service.RunTimeoutSweep()
	r := a.CheckLive()
	timeouts := clientEntriesByKind(r.Entries, constant.QueueTimeout)
	assert.Count(t, 1, timeouts)
	assert.String(
		t,
		"30 minutes since completing. Removed from roster.",
		timeouts[0].Body,
	)
}

func TestNoFalseTimeout(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "still working")
	a.CheckLive()
	s.Store.Advance(30 * time.Minute)
	s.Service.RunTimeoutSweep()
	r := a.CheckLive()
	timeouts := clientEntriesByKind(r.Entries, constant.QueueTimeout)
	assert.Count(t, 0, timeouts)
}

func TestTimeoutClearsOnCheck(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.CheckLive()
	s.Store.Advance(2 * time.Hour)
	s.Service.RunTimeoutSweep()
	first := a.CheckLive()
	timeouts := clientEntriesByKind(first.Entries, constant.QueueTimeout)
	assert.Count(t, 1, timeouts)
	second := a.CheckLive()
	timeouts = clientEntriesByKind(second.Entries, constant.QueueTimeout)
	assert.Count(t, 0, timeouts)
}

func TestTimeoutVisibleToOthers(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "going quiet")
	b.Announce(b.Name(), "watching")
	s.Store.Advance(2 * time.Hour)
	s.Service.RunTimeoutSweep()
	v := b.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "removed from roster", v)
	assert.StringContains(t, a.Name(), v)
}

package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestRelease(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "temporary work")
	b.Announce(b.Name(), "watching")
	b.CheckLive()
	a.MustCallTool(constant.Release, map[string]any{})
	r := b.CheckLive()
	releases := clientEntriesByKind(r.Entries, constant.QueueSessionRelease)
	assert.True(t, len(releases) > 0)
	assert.StringContains(t, a.Name(), releases[0].Body)
}

func TestReleaseHistoryEvent(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "observer")
	b.Announce(b.Name(), "departing")
	b.MustCallTool(constant.Release, map[string]any{})
	h := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, b.Name(), h)
	assert.StringContains(t, "left", h)
}

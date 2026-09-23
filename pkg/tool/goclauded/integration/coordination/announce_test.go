package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"testing"
)

func TestAnnounce(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "reviewing proposals")
	r := a.CheckLive()
	announces := clientEntriesByKind(r.Entries, constant.QueueSessionAnnounce)
	assert.True(t, len(announces) > 0)
	assert.StringContains(t, "reviewing proposals", announces[0].Body)
}

func TestAnnounceReannounce(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "first topic")
	a.CheckLive()
	a.Announce(a.Name(), "second topic")
	r := a.CheckLive()
	announces := clientEntriesByKind(r.Entries, constant.QueueSessionAnnounce)
	assert.True(t, len(announces) > 0)
	assert.StringContains(t, "second topic", announces[0].Body)
}

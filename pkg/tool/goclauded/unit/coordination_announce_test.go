package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/delivery_tester"
	"testing"
)

func TestAnnounceReachesSessionActivity(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "reviewing proposals")
	r := a.CheckLive()
	announces := delivery_tester.ClientKindEntries(
		r.Entries,
		constant.QueueSessionAnnounce,
	)
	assert.True(t, len(announces) > 0)
	assert.StringContains(t, "reviewing proposals", announces[0].Body)
}

func TestAnnounceReannounce(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "first topic")
	a.CheckLive()
	a.Announce(a.Name(), "second topic")
	r := a.CheckLive()
	announces := delivery_tester.ClientKindEntries(
		r.Entries,
		constant.QueueSessionAnnounce,
	)
	assert.True(t, len(announces) > 0)
	assert.StringContains(t, "second topic", announces[0].Body)
}

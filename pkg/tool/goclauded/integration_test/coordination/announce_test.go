package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func clientEntriesByKind(
	entries []client.QueueEntry,
	kind string,
) []client.QueueEntry {
	var result []client.QueueEntry

	for _, e := range entries {
		if e.Kind == kind {
			result = append(result, e)
		}
	}

	return result
}

func TestAnnounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "reviewing proposals")
	r := a.CheckLive()
	announces := clientEntriesByKind(r.Entries, constant.QueueSessionAnnounce)
	assert.True(t, len(announces) > 0)
	assert.StringContains(t, "reviewing proposals", announces[0].Body)
}

func TestAnnounceReannounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "first topic")
	a.CheckLive()
	a.Announce(a.Name(), "second topic")
	r := a.CheckLive()
	announces := clientEntriesByKind(r.Entries, constant.QueueSessionAnnounce)
	assert.True(t, len(announces) > 0)
	assert.StringContains(t, "second topic", announces[0].Body)
}

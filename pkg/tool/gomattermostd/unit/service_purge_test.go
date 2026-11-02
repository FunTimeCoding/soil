package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"path/filepath"
	"testing"
	"time"
)

func TestPurgeDropsStaleOnlyAndForgetsTheRoot(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	m.AddPost("alfa")
	m.AddPost("charlie")
	sink, k := worker_tester.NewSink(t)
	path := filepath.Join(t.TempDir(), "mattermost.db")
	s := service_tester.Reported(
		t,
		m,
		k,
		func(o *option.Mattermost) { o.LitePath = path },
		memory.New(),
	)
	store := worker_tester.NewStoreOnFile(t, path)
	subscribe(t, s, m, "alfa", "kilo", "papa")
	subscribe(t, s, m, "charlie", "kilo", "quebec")
	store.MustTouchRoot("alfa", time.Now().Add(-8*24*time.Hour))
	assert.Count(t, 0, s.Purge())
	remaining := store.MustAll()
	assert.Count(t, 1, remaining)
	assert.String(t, "charlie", remaining[0].RootIdentifier)
	assert.False(t, s.Watching("alfa"))
	assert.True(t, s.Watching("charlie"))
	notified := sink.All()
	assert.Count(t, 1, notified)
	assert.String(t, "kilo", notified[0].Callsign)
	assert.StringContains(t, "papa", notified[0].Body)
	assert.StringContains(t, "7 days", notified[0].Body)
}

func TestPurgeIsQuietWhenNothingIsStale(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	m.AddPost("alfa")
	sink, k := worker_tester.NewSink(t)
	s := service_tester.Service(t, m, k, nil)
	subscribe(t, s, m, "alfa", "kilo", "papa")
	assert.Count(t, 0, s.Purge())
	assert.True(t, s.Watching("alfa"))
	assert.Count(t, 0, sink.All())
}

func TestPurgeKeepsWatchingARootALiveSubscriberHolds(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	m.AddPost("alfa")
	sink, k := worker_tester.NewSink(t)
	path := filepath.Join(t.TempDir(), "mattermost.db")
	s := service_tester.Reported(
		t,
		m,
		k,
		func(o *option.Mattermost) { o.LitePath = path },
		memory.New(),
	)
	store := worker_tester.NewStoreOnFile(t, path)
	subscribe(t, s, m, "alfa", "kilo", "papa")
	subscribe(t, s, m, "alfa", "lima", "papa")
	store.MustTouchRoot("alfa", time.Now().Add(-8*24*time.Hour))
	subscribe(t, s, m, "alfa", "mike", "sierra")
	assert.Count(t, 0, s.Purge())
	remaining := store.MustAll()
	assert.Count(t, 1, remaining)
	assert.String(t, "mike", remaining[0].Callsign)
	assert.True(t, s.Watching("alfa"))
	assert.Count(t, 2, sink.All())
}

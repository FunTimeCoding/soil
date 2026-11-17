package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"testing"
	"time"
)

func TestPurgeDropsStaleOnlyAndForgetsRoot(t *testing.T) {
	sink, k := worker_tester.NewSink(t)
	o := service_tester.New(t, service_tester.WithConnector(k))
	o.Mattermost.AddPost("alfa")
	o.Mattermost.AddPost("charlie")
	s := o.Service()
	store := worker_tester.NewStoreOnFile(t, o.Path)
	o.Subscribe("alfa", "kilo", "papa")
	o.Subscribe("charlie", "kilo", "quebec")
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
	assert.String(t, "queued", string(*notified[0].Delivery))
}

func TestPurgeIsQuietWhenNothingIsStale(t *testing.T) {
	sink, k := worker_tester.NewSink(t)
	o := service_tester.New(t, service_tester.WithConnector(k))
	o.Mattermost.AddPost("alfa")
	s := o.Service()
	o.Subscribe("alfa", "kilo", "papa")
	assert.Count(t, 0, s.Purge())
	assert.True(t, s.Watching("alfa"))
	assert.Count(t, 0, sink.All())
}

func TestPurgeKeepsWatchingRootLiveSubscriberHolds(t *testing.T) {
	sink, k := worker_tester.NewSink(t)
	o := service_tester.New(t, service_tester.WithConnector(k))
	o.Mattermost.AddPost("alfa")
	s := o.Service()
	store := worker_tester.NewStoreOnFile(t, o.Path)
	o.Subscribe("alfa", "kilo", "papa")
	o.Subscribe("alfa", "lima", "papa")
	store.MustTouchRoot("alfa", time.Now().Add(-8*24*time.Hour))
	o.Subscribe("alfa", "mike", "sierra")
	assert.Count(t, 0, s.Purge())
	remaining := store.MustAll()
	assert.Count(t, 1, remaining)
	assert.String(t, "mike", remaining[0].Callsign)
	assert.True(t, s.Watching("alfa"))
	assert.Count(t, 2, sink.All())
}

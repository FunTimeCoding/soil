package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"testing"
	"time"
)

func TestSweeperSweepDropsStaleOnly(t *testing.T) {
	o := worker_tester.NewSweeper(t)
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Store.MustCreate(subscription.New("kilo", "charlie", "bravo", "quebec"))
	o.Store.MustTouchRoot("alfa", time.Now().Add(-8*24*time.Hour))
	o.Sweeper.Sweep()
	remaining := o.Store.MustAll()
	assert.Integer(t, 1, len(remaining))
	assert.String(t, "charlie", remaining[0].RootIdentifier)
	assert.Strings(t, []string{"alfa"}, o.Index.Forgotten())
	notified := o.Sink.All()
	assert.Integer(t, 1, len(notified))
	assert.String(t, "kilo", notified[0].Callsign)
	assert.StringContains(t, "papa", notified[0].Body)
	assert.StringContains(t, "7 days", notified[0].Body)
}

func TestSweeperSweepQuietWhenNothingStale(t *testing.T) {
	o := worker_tester.NewSweeper(t)
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Sweeper.Sweep()
	assert.Integer(t, 1, len(o.Store.MustAll()))
	assert.Integer(t, 0, len(o.Index.Forgotten()))
	assert.Integer(t, 0, len(o.Sink.All()))
}

func TestSweeperSweepKeepsIndexForLiveSubscriber(t *testing.T) {
	o := worker_tester.NewSweeper(t)
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Store.MustCreate(subscription.New("lima", "alfa", "bravo", "papa"))
	o.Store.MustTouchRoot("alfa", time.Now().Add(-8*24*time.Hour))
	o.Store.MustCreate(subscription.New("mike", "alfa", "bravo", "sierra"))
	o.Sweeper.Sweep()
	remaining := o.Store.MustAll()
	assert.Integer(t, 1, len(remaining))
	assert.String(t, "mike", remaining[0].Callsign)
	assert.Integer(t, 0, len(o.Index.Forgotten()))
	assert.Integer(t, 2, len(o.Sink.All()))
}

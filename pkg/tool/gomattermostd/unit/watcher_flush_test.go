package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"strings"
	"testing"
	"time"
)

func TestWatcherFlushNotifiesEverySubscriber(t *testing.T) {
	o := worker_tester.NewWatcher(t, time.Hour)
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Store.MustCreate(subscription.New("lima", "alfa", "bravo", ""))
	o.Watcher.Record(
		"alfa",
		event.New(
			constant.MessageEvent,
			"Foxtrot",
			"first message",
			time.Date(2026, 9, 8, 14, 32, 0, 0, time.UTC),
		),
	)
	o.Watcher.Flush("alfa")
	result := o.Sink.All()
	assert.Integer(t, 2, len(result))
	assert.String(t, "kilo", result[0].Callsign)
	assert.String(t, "[papa] 14:32 · Foxtrot: first message", result[0].Body)
	assert.String(t, "lima", result[1].Callsign)
	assert.String(t, "[alfa] 14:32 · Foxtrot: first message", result[1].Body)
}

func TestWatcherFlushCoalescesPerThread(t *testing.T) {
	o := worker_tester.NewWatcher(t, time.Hour)
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	at := time.Date(2026, 9, 8, 14, 32, 0, 0, time.UTC)
	o.Watcher.Record(
		"alfa",
		event.New(constant.MessageEvent, "Foxtrot", "one", at),
	)
	o.Watcher.Record(
		"alfa",
		event.New(constant.MessageEvent, "Foxtrot", "two", at.Add(time.Minute)),
	)
	o.Watcher.Record(
		"alfa",
		event.New(
			constant.ReactionAddedEvent,
			"Golf",
			":eyes:",
			at.Add(2*time.Minute),
		),
	)
	o.Watcher.Flush("alfa")
	result := o.Sink.All()
	assert.Integer(t, 1, len(result))
	assert.Strings(
		t,
		[]string{
			"[papa] 3 new · 14:32-14:34",
			"    Foxtrot: one",
			"    Foxtrot: two",
			"    Golf reacted :eyes:",
		},
		strings.Split(result[0].Body, "\n"),
	)
}

func TestWatcherFlushEmptyBufferSendsNothing(t *testing.T) {
	o := worker_tester.NewWatcher(t, time.Hour)
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Watcher.Flush("alfa")
	assert.Integer(t, 0, len(o.Sink.All()))
}

func TestWatcherFlushTouchesSubscription(t *testing.T) {
	o := worker_tester.NewWatcher(t, time.Hour)
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Store.MustTouchRoot("alfa", time.Now().Add(-8*24*time.Hour))
	o.Watcher.Record(
		"alfa",
		event.New(constant.MessageEvent, "Foxtrot", "one", time.Now()),
	)
	o.Watcher.Flush("alfa")
	assert.True(
		t,
		o.Store.MustByRoot("alfa")[0].LastEvent.After(
			time.Now().Add(-1*time.Minute),
		),
	)
}

func TestWatcherDebounceFlushesAfterWindow(t *testing.T) {
	o := worker_tester.NewWatcher(t, 10*time.Millisecond)
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Watcher.Record(
		"alfa",
		event.New(constant.MessageEvent, "Foxtrot", "one", time.Now()),
	)
	assert.Integer(t, 0, len(o.Sink.All()))
	deadline := time.Now().Add(5 * time.Second)

	for time.Now().Before(deadline) && len(o.Sink.All()) == 0 {
		time.Sleep(5 * time.Millisecond)
	}

	assert.Integer(t, 1, len(o.Sink.All()))
}

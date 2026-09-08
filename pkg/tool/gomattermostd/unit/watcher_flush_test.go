package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"strings"
	"testing"
	"time"
)

func newWatcher(
	t *testing.T,
	window time.Duration,
) (*watcher.Watcher, *store.Store, *notifySink) {
	t.Helper()
	sink, c := newNotifyClient(t)
	s := store.New(lite.NewMemory())
	r := memory.New()

	return watcher.New(
		mock_client.New("selfuser"),
		s,
		notifier.New(c, "mattermost", r),
		logger.New(context.Background()),
		r,
		window,
	), s, sink
}

func TestWatcherFlushNotifiesEverySubscriber(t *testing.T) {
	w, s, sink := newWatcher(t, time.Hour)
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	s.MustCreate(subscription.New("lima", "alfa", "bravo", ""))
	w.Record(
		"alfa",
		event.New(
			constant.MessageEvent,
			"Foxtrot",
			"first message",
			time.Date(2026, 9, 8, 14, 32, 0, 0, time.UTC),
		),
	)
	w.Flush("alfa")
	result := sink.all()
	assert.Integer(t, 2, len(result))
	assert.String(t, "kilo", result[0].Callsign)
	assert.String(t, "[papa] 14:32 · Foxtrot: first message", result[0].Body)
	assert.String(t, "lima", result[1].Callsign)
	assert.String(t, "[alfa] 14:32 · Foxtrot: first message", result[1].Body)
}

func TestWatcherFlushCoalescesPerThread(t *testing.T) {
	w, s, sink := newWatcher(t, time.Hour)
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	at := time.Date(2026, 9, 8, 14, 32, 0, 0, time.UTC)
	w.Record("alfa", event.New(constant.MessageEvent, "Foxtrot", "one", at))
	w.Record(
		"alfa",
		event.New(constant.MessageEvent, "Foxtrot", "two", at.Add(time.Minute)),
	)
	w.Record(
		"alfa",
		event.New(
			constant.ReactionAddedEvent,
			"Golf",
			":eyes:",
			at.Add(2*time.Minute),
		),
	)
	w.Flush("alfa")
	result := sink.all()
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
	w, s, sink := newWatcher(t, time.Hour)
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	w.Flush("alfa")
	assert.Integer(t, 0, len(sink.all()))
}

func TestWatcherFlushTouchesSubscription(t *testing.T) {
	w, s, _ := newWatcher(t, time.Hour)
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	s.MustTouchRoot("alfa", time.Now().Add(-8*24*time.Hour))
	w.Record(
		"alfa",
		event.New(constant.MessageEvent, "Foxtrot", "one", time.Now()),
	)
	w.Flush("alfa")
	assert.True(
		t,
		s.MustByRoot("alfa")[0].LastEvent.After(time.Now().Add(-1*time.Minute)),
	)
}

func TestWatcherDebounceFlushesAfterWindow(t *testing.T) {
	w, s, sink := newWatcher(t, 10*time.Millisecond)
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	w.Record(
		"alfa",
		event.New(constant.MessageEvent, "Foxtrot", "one", time.Now()),
	)
	assert.Integer(t, 0, len(sink.all()))
	deadline := time.Now().Add(5 * time.Second)

	for time.Now().Before(deadline) && len(sink.all()) == 0 {
		time.Sleep(5 * time.Millisecond)
	}

	assert.Integer(t, 1, len(sink.all()))
}

package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/mock_indexer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/sweeper"
	"testing"
	"time"
)

func newSweeper(
	t *testing.T,
) (*sweeper.Sweeper, *store.Store, *mock_indexer.Indexer, *notifySink) {
	t.Helper()
	sink, c := newNotifyClient(t)
	s := store.New(lite.NewMemory())
	r := memory.New()
	index := mock_indexer.New()

	return sweeper.New(
		s,
		notifier.New(c, "mattermost", r),
		index,
		logger.New(context.Background()),
		r,
		constant.PurgeWindow,
		constant.PurgeInterval,
	), s, index, sink
}

func TestSweeperSweepDropsStaleOnly(t *testing.T) {
	w, s, index, sink := newSweeper(t)
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	s.MustCreate(subscription.New("kilo", "charlie", "bravo", "quebec"))
	s.MustTouchRoot("alfa", time.Now().Add(-8*24*time.Hour))
	w.Sweep()
	remaining := s.MustAll()
	assert.Integer(t, 1, len(remaining))
	assert.String(t, "charlie", remaining[0].RootIdentifier)
	assert.Strings(t, []string{"alfa"}, index.Forgotten())
	notified := sink.all()
	assert.Integer(t, 1, len(notified))
	assert.String(t, "kilo", notified[0].Callsign)
	assert.StringContains(t, "papa", notified[0].Body)
	assert.StringContains(t, "7 days", notified[0].Body)
}

func TestSweeperSweepQuietWhenNothingStale(t *testing.T) {
	w, s, index, sink := newSweeper(t)
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	w.Sweep()
	assert.Integer(t, 1, len(s.MustAll()))
	assert.Integer(t, 0, len(index.Forgotten()))
	assert.Integer(t, 0, len(sink.all()))
}

func TestSweeperSweepKeepsIndexForLiveSubscriber(t *testing.T) {
	w, s, index, sink := newSweeper(t)
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	s.MustCreate(subscription.New("lima", "alfa", "bravo", "papa"))
	s.MustTouchRoot("alfa", time.Now().Add(-8*24*time.Hour))
	s.MustCreate(subscription.New("mike", "alfa", "bravo", "sierra"))
	w.Sweep()
	remaining := s.MustAll()
	assert.Integer(t, 1, len(remaining))
	assert.String(t, "mike", remaining[0].Callsign)
	assert.Integer(t, 0, len(index.Forgotten()))
	assert.Integer(t, 2, len(sink.all()))
}

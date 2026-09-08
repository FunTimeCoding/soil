package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"testing"
	"time"
)

func TestStoreSubscriptionRoundTrip(t *testing.T) {
	s := store.New(lite.NewMemory())
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	s.MustCreate(subscription.New("lima", "charlie", "bravo", ""))
	own := s.MustByCallsign("kilo")
	assert.Integer(t, 1, len(own))
	assert.String(t, "alfa", own[0].RootIdentifier)
	assert.String(t, "papa", own[0].Alias)
	assert.String(t, "bravo", own[0].ChannelIdentifier)
	root := s.MustByRoot("charlie")
	assert.Integer(t, 1, len(root))
	assert.String(t, "lima", root[0].Callsign)
	assert.Integer(t, 2, len(s.MustAll()))
}

func TestStoreSubscriptionDelete(t *testing.T) {
	s := store.New(lite.NewMemory())
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", ""))
	s.MustCreate(subscription.New("lima", "alfa", "bravo", ""))
	assert.Integer(t, 1, s.MustDelete("kilo", "alfa"))
	remaining := s.MustByRoot("alfa")
	assert.Integer(t, 1, len(remaining))
	assert.String(t, "lima", remaining[0].Callsign)
	assert.Integer(t, 0, s.MustDelete("kilo", "alfa"))
}

func TestStoreSubscriptionTouchRoot(t *testing.T) {
	s := store.New(lite.NewMemory())
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", ""))
	s.MustCreate(subscription.New("lima", "alfa", "bravo", ""))
	s.MustCreate(subscription.New("kilo", "charlie", "bravo", ""))
	stale := time.Now().Add(-9 * 24 * time.Hour)
	s.MustTouchRoot("alfa", stale)

	for _, v := range s.MustByRoot("alfa") {
		assert.True(t, v.LastEvent.Before(time.Now().Add(-8*24*time.Hour)))
	}

	assert.True(
		t,
		s.MustByRoot("charlie")[0].LastEvent.After(
			time.Now().Add(-1*time.Hour),
		),
	)
}

func TestStoreSubscriptionPrune(t *testing.T) {
	s := store.New(lite.NewMemory())
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "romeo"))
	s.MustCreate(subscription.New("kilo", "charlie", "bravo", "quebec"))
	s.MustTouchRoot("alfa", time.Now().Add(-8*24*time.Hour))
	pruned := s.MustPrune(time.Now().Add(-7 * 24 * time.Hour))
	assert.Integer(t, 1, len(pruned))
	assert.String(t, "alfa", pruned[0].RootIdentifier)
	assert.String(t, "romeo", pruned[0].Alias)
	assert.String(t, "kilo", pruned[0].Callsign)
	remaining := s.MustAll()
	assert.Integer(t, 1, len(remaining))
	assert.String(t, "charlie", remaining[0].RootIdentifier)
}

func TestStoreSubscriptionPruneEmpty(t *testing.T) {
	s := store.New(lite.NewMemory())
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", ""))
	assert.Integer(t, 0, len(s.MustPrune(time.Now().Add(-7*24*time.Hour))))
	assert.Integer(t, 1, len(s.MustAll()))
}

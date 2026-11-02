package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/cross_service_tester"
	"testing"
	"time"
)

func TestTheFlushNotifiesEverySubscriberAndTouchesTheRoot(t *testing.T) {
	asked := make(chan struct{}, 1)
	release := make(chan struct{})
	close(release)
	r := cross_service_tester.New(t, heldAuthorUpstream(asked, release), 0)
	r.Goclauded.Store.HoldCallsign("kilo-session", "kilo")
	r.Goclauded.Store.HoldCallsign("lima-session", "lima")
	r.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	r.Store.MustCreate(subscription.New("lima", "alfa", "bravo", ""))
	r.Store.MustTouchRoot("alfa", time.Now().Add(-8*24*time.Hour))
	r.Start()
	pushReply(r, asked)
	eventually(
		t,
		func() bool {
			return r.Store.MustByRoot("alfa")[0].LastEvent.After(
				time.Now().Add(-time.Minute),
			)
		},
	)
	kilo := r.Goclauded.Store.PendingNotifications("kilo-session")
	assert.Count(t, 1, kilo)
	assert.StringContains(t, "[papa]", kilo[0].Body)
	lima := r.Goclauded.Store.PendingNotifications("lima-session")
	assert.Count(t, 1, lima)
	assert.StringContains(t, "[alfa]", lima[0].Body)
}

func TestARootInsideItsWindowWaitsForCloseToFlush(t *testing.T) {
	asked := make(chan struct{}, 1)
	release := make(chan struct{})
	close(release)
	r := cross_service_tester.New(
		t,
		heldAuthorUpstream(asked, release),
		time.Hour,
	)
	r.Goclauded.Store.HoldCallsign("kilo-session", "kilo")
	r.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	r.Start()
	pushReply(r, asked)
	assert.Count(t, 0, r.Service.FlushDue())
	assert.Count(t, 0, r.Goclauded.Store.PendingNotifications("kilo-session"))
	r.Service.Close()
	assert.Count(t, 1, r.Goclauded.Store.PendingNotifications("kilo-session"))
}

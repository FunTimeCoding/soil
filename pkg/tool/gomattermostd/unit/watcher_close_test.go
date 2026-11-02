package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/cross_service_tester"
	"github.com/mattermost/mattermost/server/public/model"
	"sync"
	"testing"
	"time"
)

func TestCloseFlushesAnEventStillBeingDispatched(t *testing.T) {
	asked := make(chan struct{}, 1)
	release := make(chan struct{})
	var releasing sync.Once
	r := cross_service_tester.New(
		t,
		heldAuthorUpstream(asked, release),
		time.Hour,
	)
	t.Cleanup(func() { releasing.Do(func() { close(release) }) })
	r.Goclauded.Store.HoldCallsign("kilo-session", "kilo")
	r.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	r.Start()
	r.Upstream.Push(
		postEvent(
			&model.Post{
				Id:       "reply",
				RootId:   "alfa",
				UserId:   "foxtrot",
				Message:  "late message",
				CreateAt: time.Now().UnixMilli(),
			},
		),
	)

	select {
	case <-asked:
	case <-time.After(5 * time.Second):
		t.Fatal("the event never reached dispatch")
	}

	closed := make(chan struct{})
	go func() {
		r.Service.Close()
		close(closed)
	}()

	select {
	case <-closed:
		t.Fatal("close returned with an event in dispatch")
	case <-time.After(100 * time.Millisecond):
	}

	releasing.Do(func() { close(release) })

	select {
	case <-closed:
	case <-time.After(5 * time.Second):
		t.Fatal("close never returned")
	}

	result := r.Goclauded.Store.PendingNotifications("kilo-session")
	assert.Count(t, 1, result)
	assert.StringContains(t, "late message", result[0].Body)
}

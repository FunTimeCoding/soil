package worker

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/notification"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/integration/cross_service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
	"time"
)

func waitForNotification(
	t *testing.T,
	r *cross_service_tester.Tester,
) []notification.Notification {
	t.Helper()
	var result []notification.Notification
	deadline := time.Now().Add(constant.MattermostSocketWait)

	for time.Now().Before(deadline) {
		result = r.Goclauded.Store.PendingNotifications("kilo")

		if len(result) > 0 {
			return result
		}

		time.Sleep(5 * time.Millisecond)
	}

	return result
}

func TestWatcherResumesAfterDroppedConnection(t *testing.T) {
	r := cross_service_tester.New(t, upstream, 10*time.Millisecond)
	r.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	r.Watcher.Start()
	r.Upstream.Push(
		postEvent(
			&model.Post{
				Id:       "reply",
				RootId:   "alfa",
				UserId:   "foxtrot",
				Message:  "first message",
				CreateAt: time.Now().UnixMilli(),
			},
		),
	)
	assert.Integer(t, 1, len(waitForNotification(t, r)))
	r.Upstream.Drop()
	deadline := time.Now().Add(30 * time.Second)

	for time.Now().Before(deadline) && r.Upstream.Accepted() < 2 {
		time.Sleep(20 * time.Millisecond)
	}

	assert.Integer(t, 2, r.Upstream.Accepted())
	r.Upstream.Push(
		postEvent(
			&model.Post{
				Id:       "second",
				RootId:   "alfa",
				UserId:   "foxtrot",
				Message:  "second message",
				CreateAt: time.Now().UnixMilli(),
			},
		),
	)
	result := waitForNotification(t, r)
	assert.Integer(t, 1, len(result))
	assert.StringContains(t, "second message", result[0].Body)
	reported := r.Reporter.Events()
	assert.Integer(t, 1, len(reported))
	assert.StringContains(
		t,
		"websocket event channel closed, reconnecting",
		reported[0].Error.Error(),
	)
}

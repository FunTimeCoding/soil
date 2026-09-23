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

func TestSocketEventBecomesNotification(t *testing.T) {
	r := cross_service_tester.New(t, upstream, 10*time.Millisecond)
	r.Goclauded.Store.HoldCallsign("kilo-session", "kilo")
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
	var result []notification.Notification
	deadline := time.Now().Add(constant.MattermostSocketWait)

	for time.Now().Before(deadline) {
		result = r.Goclauded.Store.PendingNotifications("kilo-session")

		if len(result) > 0 {
			break
		}

		time.Sleep(5 * time.Millisecond)
	}

	assert.Integer(t, 0, len(r.Reporter.Events()))
	assert.Integer(t, 1, len(result))
	assert.String(t, "mattermost", result[0].Source)
	assert.StringContains(t, "[papa]", result[0].Body)
	assert.StringContains(t, "Foxtrot: first message", result[0].Body)
}

func TestSocketEventInUnwatchedThreadIgnored(t *testing.T) {
	r := cross_service_tester.New(t, upstream, 10*time.Millisecond)
	r.Goclauded.Store.HoldCallsign("kilo-session", "kilo")
	r.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	r.Watcher.Start()
	r.Upstream.Push(
		postEvent(
			&model.Post{
				Id:       "reply",
				RootId:   "elsewhere",
				UserId:   "foxtrot",
				Message:  "not our thread",
				CreateAt: time.Now().UnixMilli(),
			},
		),
	)
	time.Sleep(200 * time.Millisecond)
	assert.Integer(
		t,
		0,
		len(r.Goclauded.Store.PendingNotifications("kilo-session")),
	)
}

package worker

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/notification"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/integration/cross_service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
	"testing"
	"time"
)

func upstream(m *http.ServeMux) {
	m.HandleFunc(
		"/api/v4/users/me",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			web.Encode(w, &model.User{Id: "self", Username: "assistant"})
		},
	)
	m.HandleFunc(
		"/api/v4/users/foxtrot",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			web.Encode(w, &model.User{Id: "foxtrot", Username: "Foxtrot"})
		},
	)
	m.HandleFunc(
		"/api/v4/posts/alfa",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			web.Encode(w, &model.Post{Id: "alfa", ChannelId: "bravo"})
		},
	)
}

func postEvent(p *model.Post) *model.WebSocketEvent {
	v := &model.WebSocketEvent{}
	v = v.SetEvent(model.WebsocketEventPosted)

	return v.SetData(
		map[string]any{constant.MattermostPostField: notation.Encode(p, false)},
	)
}

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

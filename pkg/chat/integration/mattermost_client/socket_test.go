package mattermost_client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/integration/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
	"testing"
	"time"
)

func TestSocketDeliversPushedEvent(t *testing.T) {
	r := mattermost_client_tester.New(t, func(_ *http.ServeMux) {})
	w := r.Client.WebSocket()
	w.Listen()
	defer w.Close()
	v := &model.WebSocketEvent{}
	v = v.SetEvent(model.WebsocketEventPosted)
	r.Push(
		v.SetData(
			map[string]any{
				constant.MattermostPostField: notation.Encode(
					&model.Post{
						Id:      "bravo",
						RootId:  "alfa",
						UserId:  "delta",
						Message: "pushed from the tester",
					},
					false,
				),
			},
		),
	)

	select {
	case received := <-w.EventChannel:
		assert.String(t, "posted", string(received.EventType()))
		decoded := post.Decode(received)
		assert.String(t, "bravo", decoded.Id)
		assert.String(t, "alfa", decoded.RootId)
		assert.String(t, "pushed from the tester", decoded.Message)
	case <-time.After(constant.MattermostSocketWait):
		t.Fatal("no event received")
	}
}

package mattermost_client_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/gorilla/websocket"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (t *Tester) Push(v *model.WebSocketEvent) {
	t.t.Helper()

	select {
	case <-t.ready:
	case <-time.After(constant.MattermostSocketWait):
		t.t.Fatal("no websocket connection")
	}

	b, e := v.ToJSON()
	assert.FatalOnError(t.t, e)
	t.mutex.Lock()
	defer t.mutex.Unlock()
	assert.FatalOnError(
		t.t,
		t.connection.WriteMessage(websocket.TextMessage, b),
	)
}

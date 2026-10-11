package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"net/http"
	"testing"
	"time"
)

func TestSocketPingsTheServerWhileIdle(t *testing.T) {
	r := mattermost_client_tester.New(
		t,
		func(_ *http.ServeMux) {},
		mattermost.WithKeepAlive(20*time.Millisecond),
	)
	w := r.Client.WebSocket()
	w.Listen()
	defer w.Close()
	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) && r.Pings() < 3 {
		time.Sleep(10 * time.Millisecond)
	}

	assert.True(t, r.Pings() >= 3)
}

func TestRefreshSocketMovesTheKeepAlive(t *testing.T) {
	r := mattermost_client_tester.New(
		t,
		func(_ *http.ServeMux) {},
		mattermost.WithKeepAlive(20*time.Millisecond),
	)
	before := r.Client.WebSocket()
	defer before.Close()
	before.Listen()
	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) && r.Pings() < 2 {
		time.Sleep(10 * time.Millisecond)
	}

	assert.True(t, r.Pings() >= 2)
	assert.Nil(t, r.Client.RefreshSocket())
	after := r.Client.WebSocket()
	after.Listen()
	defer after.Close()
	assert.Integer(t, 2, r.Accepted())
	seen := r.Pings()
	deadline = time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) && r.Pings() < seen+3 {
		time.Sleep(10 * time.Millisecond)
	}

	assert.True(t, r.Pings() >= seen+3)
}

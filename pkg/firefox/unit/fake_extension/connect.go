package fake_extension

import (
	"github.com/coder/websocket"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/firefox/types/message"
	"strings"
	"testing"
)

func Connect(
	t *testing.T,
	address string,
) *Extension {
	t.Helper()
	c, _, e := websocket.Dial(
		t.Context(),
		strings.Replace(address, "http://", "ws://", 1),
		nil,
	)
	assert.FatalOnError(t, e)
	result := &Extension{connection: c, replies: map[string]*message.Reply{}}
	t.Cleanup(result.Disconnect)
	go result.serve(t)

	return result
}

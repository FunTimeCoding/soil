package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/assistant/unit/assistant_server"
	"github.com/funtimecoding/soil/pkg/assistant/unit/assistant_tester"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"testing"
)

func TestStartingHomeAssistantIsWaitedOut(t *testing.T) {
	f := assistant_server.New(t)
	f.Refuse(2)
	r := memory.New()
	received := make(chan *message.Message, 8)
	c := assistant_tester.Listener(f, r, received)
	c.Start()
	defer c.Close()
	assert.Signalled(t, f.Subscribed())
	f.Emit("state_changed")
	assert.String(t, "state_changed", assert.Signalled(t, received).Event.Type)
	assert.Count(t, 0, r.Events())
}

func TestRestartedHomeAssistantIsSubscribedAgain(t *testing.T) {
	f := assistant_server.New(t)
	r := memory.New()
	received := make(chan *message.Message, 8)
	c := assistant_tester.Listener(f, r, received)
	c.Start()
	defer c.Close()
	assert.Signalled(t, f.Subscribed())
	f.Emit("state_changed")
	assert.Signalled(t, received)
	f.Restart()
	assert.Signalled(t, f.Subscribed())
	f.Emit("call_service")
	assert.String(t, "call_service", assert.Signalled(t, received).Event.Type)
	assert.Integer(t, 2, f.Subscriptions())
	assert.Count(t, 0, r.Events())
}

func TestRefusedTokenIsReportedOnce(t *testing.T) {
	f := assistant_server.New(t)
	f.RefuseToken()
	r := memory.New()
	received := make(chan *message.Message, 8)
	c := assistant_tester.Listener(f, r, received)
	c.Start()
	defer c.Close()
	assert.Signalled(t, f.Dialed())
	assert.Signalled(t, f.Dialed())
	assert.Signalled(t, f.Dialed())
	assert.Integer(t, 0, f.Subscriptions())
	assert.Count(t, 1, r.Events())
}

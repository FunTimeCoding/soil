package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/assistant/unit/fake_assistant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"testing"
)

func TestAStartingHomeAssistantIsWaitedOut(t *testing.T) {
	f := fake_assistant.New(t)
	f.Refuse(2)
	r := memory.New()
	received := make(chan *message.Message, 8)
	c := listener(f, r, received)
	c.Start()
	defer c.Close()
	await(t, f.Subscribed())
	f.Emit("state_changed")
	assert.String(t, "state_changed", await(t, received).Event.Type)
	assert.Count(t, 0, r.Events())
}

func TestARestartedHomeAssistantIsSubscribedAgain(t *testing.T) {
	f := fake_assistant.New(t)
	r := memory.New()
	received := make(chan *message.Message, 8)
	c := listener(f, r, received)
	c.Start()
	defer c.Close()
	await(t, f.Subscribed())
	f.Emit("state_changed")
	await(t, received)
	f.Restart()
	await(t, f.Subscribed())
	f.Emit("call_service")
	assert.String(t, "call_service", await(t, received).Event.Type)
	assert.Integer(t, 2, f.Subscriptions())
	assert.Count(t, 0, r.Events())
}

func TestARefusedTokenIsReportedOnce(t *testing.T) {
	f := fake_assistant.New(t)
	f.RefuseToken()
	r := memory.New()
	received := make(chan *message.Message, 8)
	c := listener(f, r, received)
	c.Start()
	defer c.Close()
	await(t, f.Dialed())
	await(t, f.Dialed())
	await(t, f.Dialed())
	assert.Integer(t, 0, f.Subscriptions())
	assert.Count(t, 1, r.Events())
}

package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assistant/unit/fake_assistant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"sync/atomic"
	"testing"
	"time"
)

func TestAStartingHomeAssistantIsWaitedOut(t *testing.T) {
	f := fake_assistant.New(t)
	f.Refuse(2)
	r := memory.New()
	var received atomic.Int32
	c := listener(f, r, &received)
	c.Start()
	defer c.Close()
	time.Sleep(100 * time.Millisecond)
	assert.Integer(t, 1, f.Subscriptions())
	f.Emit("state_changed")
	time.Sleep(50 * time.Millisecond)
	assert.Integer(t, 1, int(received.Load()))
	assert.Count(t, 0, r.Events())
}

func TestARestartedHomeAssistantIsSubscribedAgain(t *testing.T) {
	f := fake_assistant.New(t)
	r := memory.New()
	var received atomic.Int32
	c := listener(f, r, &received)
	c.Start()
	defer c.Close()
	time.Sleep(50 * time.Millisecond)
	f.Emit("state_changed")
	f.Restart()
	time.Sleep(100 * time.Millisecond)
	assert.Integer(t, 2, f.Subscriptions())
	f.Emit("state_changed")
	time.Sleep(50 * time.Millisecond)
	assert.Integer(t, 2, int(received.Load()))
	assert.Count(t, 0, r.Events())
}

func TestARefusedTokenIsReportedOnce(t *testing.T) {
	f := fake_assistant.New(t)
	f.RefuseToken()
	r := memory.New()
	var received atomic.Int32
	c := listener(f, r, &received)
	c.Start()
	defer c.Close()
	time.Sleep(100 * time.Millisecond)
	assert.Integer(t, 0, f.Subscriptions())
	assert.Count(t, 1, r.Events())
}

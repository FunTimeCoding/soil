package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/turn"
	"testing"
	"time"
)

func TestATurnIsImmediateWhenIdle(t *testing.T) {
	s := turn.New(time.Second)
	g, e := s.Take(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, g.Ahead)
	g.Release()
}

func TestTurnsRotateBetweenSessions(t *testing.T) {
	s := turn.New(time.Minute)
	holder, e := s.Take(t.Context(), "charlie")
	assert.FatalOnError(t, e)
	granted := make(chan string, 3)
	queueTurn(t, s, "alfa", "alfa one", granted)
	queueTurn(t, s, "alfa", "alfa two", granted)
	queueTurn(t, s, "bravo", "bravo one", granted)
	holder.Release()
	var order []string

	for range 3 {
		order = append(order, <-granted)
	}

	assert.Strings(t, []string{"alfa one", "bravo one", "alfa two"}, order)
}

func TestABoundedWaitAnswersBusy(t *testing.T) {
	s := turn.New(10 * time.Millisecond)
	holder, e := s.Take(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	defer holder.Release()
	_, f := s.Take(t.Context(), "bravo")
	assert.String(t, "busy, 1 calls ahead", f.Error())
	assert.Integer(t, 0, s.Waiting())
}

func TestACancelledWaiterLeavesTheQueue(t *testing.T) {
	s := turn.New(time.Minute)
	holder, e := s.Take(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	x, cancel := context.WithCancel(t.Context())
	done := make(chan error, 1)
	go func() {
		_, f := s.Take(x, "bravo")
		done <- f
	}()
	waitForQueue(t, s, 1)
	cancel()
	assert.String(t, "context canceled", (<-done).Error())
	assert.Integer(t, 0, s.Waiting())
	holder.Release()
	g, f := s.Take(t.Context(), "charlie")
	assert.FatalOnError(t, f)
	assert.Integer(t, 0, g.Ahead)
	g.Release()
}

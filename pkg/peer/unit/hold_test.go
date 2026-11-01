package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/hold"
	"github.com/funtimecoding/soil/pkg/peer/unit/held_peer"
	"testing"
	"time"
)

func TestAHeldConnectionReopensAfterItDrops(t *testing.T) {
	p := held_peer.New()
	r := memory.New()
	h := hold.New("lima", p, logger.New(context.Background()), r).WithBackoff(
		time.Millisecond,
		4*time.Millisecond,
	)
	h.Start()
	defer h.Close()
	time.Sleep(20 * time.Millisecond)
	assert.Integer(t, 1, p.Opens())
	p.End(errors.New("event channel closed"))
	time.Sleep(50 * time.Millisecond)
	assert.Integer(t, 2, p.Opens())
	assert.Integer(t, 1, p.Closes())
	assert.Count(t, 0, r.Events())
}

func TestAHeldConnectionWaitsOutAnOutage(t *testing.T) {
	p := held_peer.New()
	p.Refuse(refusedConnection(), refusedConnection(), upstreamUnavailable())
	r := memory.New()
	h := hold.New("mike", p, logger.New(context.Background()), r).WithBackoff(
		time.Millisecond,
		4*time.Millisecond,
	)
	h.Start()
	defer h.Close()
	time.Sleep(100 * time.Millisecond)
	assert.Integer(t, 4, p.Opens())
	assert.Count(t, 0, r.Events())
}

func TestAHeldConnectionReportsARefusalOnce(t *testing.T) {
	p := held_peer.New()
	p.Refuse(
		errors.New("authentication invalid"),
		errors.New("authentication invalid"),
		errors.New("authentication invalid"),
	)
	r := memory.New()
	h := hold.New(
		"november",
		p,
		logger.New(context.Background()),
		r,
	).WithBackoff(time.Millisecond, 4*time.Millisecond)
	h.Start()
	defer h.Close()
	time.Sleep(100 * time.Millisecond)
	assert.Integer(t, 4, p.Opens())
	assert.Count(t, 1, r.Events())
}

func TestClosingAHeldConnectionEndsItWithoutReopening(t *testing.T) {
	p := held_peer.New()
	h := hold.New(
		"papa",
		p,
		logger.New(context.Background()),
		memory.New(),
	).WithBackoff(time.Millisecond, 4*time.Millisecond)
	h.Start()
	time.Sleep(20 * time.Millisecond)
	h.Close()
	time.Sleep(50 * time.Millisecond)
	assert.Integer(t, 1, p.Opens())
	assert.Integer(t, 2, p.Closes())
}

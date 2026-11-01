package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/hold"
	"github.com/funtimecoding/soil/pkg/peer/hold/adapter"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
	"github.com/funtimecoding/soil/pkg/peer/unit/held_peer"
	"sync/atomic"
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

func TestAHeldConnectionReportsWhereItIsTold(t *testing.T) {
	p := held_peer.New()
	p.Refuse(errors.New("authentication invalid"), refusedConnection())
	r := memory.New()
	var crossed atomic.Int32
	var refused atomic.Int32
	h := hold.New(
		"quebec",
		p,
		logger.New(context.Background()),
		r,
	).WithBackoff(time.Millisecond, 4*time.Millisecond).WithReport(
		func(*edge.Edge) { crossed.Add(1) },
		func(error) { refused.Add(1) },
	)
	h.Start()
	defer h.Close()
	time.Sleep(50 * time.Millisecond)
	assert.Integer(t, 1, int(refused.Load()))
	assert.Integer(t, 2, int(crossed.Load()))
	assert.Count(t, 0, r.Events())
}

func TestAHeldConnectionIsUnavailableUntilItServes(t *testing.T) {
	p := held_peer.New()
	p.Refuse(refusedConnection())
	h := hold.New(
		"romeo",
		p,
		logger.New(context.Background()),
		memory.New(),
	).WithBackoff(20*time.Millisecond, 40*time.Millisecond)
	e := h.Unavailable()
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "not connected", e.Error())
	h.Start()
	defer h.Close()
	time.Sleep(10 * time.Millisecond)
	assert.StringContains(t, "romeo unreachable since", h.Unavailable().Error())
	time.Sleep(50 * time.Millisecond)
	assert.Nil(t, h.Unavailable())
}

func TestAnAdapterHoldsThreeFunctions(t *testing.T) {
	var opened, closed atomic.Int32
	end := make(chan error, 1)
	h := hold.New(
		"sierra",
		adapter.New(
			func(context.Context) error {
				opened.Add(1)

				return nil
			},
			func(x context.Context) error {
				select {
				case e := <-end:
					return e
				case <-x.Done():
					return nil
				}
			},
			func() { closed.Add(1) },
		),
		logger.New(context.Background()),
		memory.New(),
	).WithBackoff(time.Millisecond, 4*time.Millisecond)
	h.Start()
	defer h.Close()
	time.Sleep(20 * time.Millisecond)
	end <- errors.New("device disconnected")
	time.Sleep(50 * time.Millisecond)
	assert.Integer(t, 2, int(opened.Load()))
	assert.Integer(t, 1, int(closed.Load()))
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
	<-h.Done()
	time.Sleep(20 * time.Millisecond)
	assert.Integer(t, 1, p.Opens())
	assert.Integer(t, 2, p.Closes())
}

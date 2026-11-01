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
	await(t, p.Served())
	p.End(errors.New("event channel closed"))
	await(t, p.Served())
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
	await(t, p.Served())
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
	await(t, p.Served())
	assert.Integer(t, 4, p.Opens())
	assert.Count(t, 1, r.Events())
}

func TestAHeldConnectionReportsWhereItIsTold(t *testing.T) {
	p := held_peer.New()
	p.Refuse(errors.New("authentication invalid"), refusedConnection())
	r := memory.New()
	crossed := make(chan *edge.Edge, 4)
	refused := make(chan error, 4)
	h := hold.New(
		"quebec",
		p,
		logger.New(context.Background()),
		r,
	).WithBackoff(time.Millisecond, 4*time.Millisecond).WithReport(
		func(x *edge.Edge) { crossed <- x },
		func(e error) { refused <- e },
	)
	h.Start()
	defer h.Close()
	await(t, p.Served())
	assert.Count(t, 1, refused)
	assert.Count(t, 2, crossed)
	assert.True(t, (<-crossed).Down)
	assert.False(t, (<-crossed).Down)
	assert.Count(t, 0, r.Events())
}

func TestAHeldConnectionIsUnavailableUntilItServes(t *testing.T) {
	p := held_peer.New()
	p.Refuse(refusedConnection())
	crossed := make(chan *edge.Edge, 4)
	h := hold.New(
		"romeo",
		p,
		logger.New(context.Background()),
		memory.New(),
	).WithBackoff(time.Hour, time.Hour).WithReport(
		func(x *edge.Edge) { crossed <- x },
		func(error) {},
	)
	e := h.Unavailable()
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "not connected", e.Error())
	h.Start()
	defer h.Close()
	assert.True(t, await(t, crossed).Down)
	assert.StringContains(t, "romeo unreachable since", h.Unavailable().Error())
}

func TestAServingHoldIsAvailable(t *testing.T) {
	p := held_peer.New()
	h := hold.New("tango", p, logger.New(context.Background()), memory.New())
	h.Start()
	defer h.Close()
	await(t, p.Served())
	assert.Nil(t, h.Unavailable())
}

func TestAnAdapterHoldsThreeFunctions(t *testing.T) {
	opened := make(chan struct{}, 4)
	closed := make(chan struct{}, 4)
	end := make(chan error, 1)
	h := hold.New(
		"sierra",
		adapter.New(
			func(context.Context) error {
				opened <- struct{}{}

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
			func() { closed <- struct{}{} },
		),
		logger.New(context.Background()),
		memory.New(),
	).WithBackoff(time.Millisecond, 4*time.Millisecond)
	h.Start()
	defer h.Close()
	await(t, opened)
	end <- errors.New("device disconnected")
	await(t, closed)
	await(t, opened)
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
	await(t, p.Served())
	h.Close()
	await(t, h.Done())
	assert.Integer(t, 1, p.Opens())
	assert.Integer(t, 2, p.Closes())
}

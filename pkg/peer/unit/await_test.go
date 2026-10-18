package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"testing"
	"time"
)

func TestAPeerIsAwaitedUntilItAnswers(t *testing.T) {
	calls := 0
	e := peer.Await(
		t.Context(),
		logger.New(context.Background()),
		"alpha",
		"http://server.test:8080",
		func(context.Context) error {
			calls++

			if calls < 2 {
				return errors.New("unanswered")
			}

			return nil
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, calls)
}

func TestAwaitingAPeerStopsWhenTheCallerStops(t *testing.T) {
	x, cancel := context.WithCancel(t.Context())
	time.AfterFunc(10*time.Millisecond, cancel)
	start := time.Now()
	e := peer.Await(
		x,
		logger.New(context.Background()),
		"bravo",
		"http://server.test:8080",
		func(context.Context) error {
			return errors.New("unanswered")
		},
	)
	assert.True(t, time.Since(start) < time.Second)
	assert.True(t, errors.Is(e, context.Canceled))
	assert.StringContains(t, "bravo not reachable", e.Error())
}

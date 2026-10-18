package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/web/peer"
	"testing"
)

func TestAPeerIsAwaitedUntilItAnswers(t *testing.T) {
	calls := 0
	e := peer.Await(
		logger.New(context.Background()),
		"alpha",
		"http://server.test:8080",
		func() error {
			calls++

			if calls < 2 {
				return errors.New("down")
			}

			return nil
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, calls)
}

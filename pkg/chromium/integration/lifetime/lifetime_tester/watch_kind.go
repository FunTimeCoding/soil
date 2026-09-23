//go:build browser

package lifetime_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/event"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"testing"
)

func WatchKind(
	t *testing.T,
	s *base.Stack,
	kind constant.EventKind,
) chan *event.Event {
	t.Helper()
	result := make(chan *event.Event, constant.FixtureWatchBuffer)
	assert.FatalOnError(
		t,
		s.Client.Watch(
			func(e *event.Event) {
				if e.Kind != kind {
					return
				}

				select {
				case result <- e:
				default:
				}
			},
		),
	)

	return result
}

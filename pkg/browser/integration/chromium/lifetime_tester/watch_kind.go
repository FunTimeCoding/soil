//go:build browser

package lifetime_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/browser/chromium/event"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/integration/chromium/base"
)

func (o *Tester) WatchKind(
	s *base.Stack,
	kind constant.ChromiumEventKind,
) chan *event.Event {
	o.t.Helper()
	result := make(chan *event.Event, constant.ChromiumFixtureWatchBuffer)
	assert.FatalOnError(
		o.t,
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

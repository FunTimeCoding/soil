//go:build browser

package lifetime_tester

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/event"
	"time"
)

func (o *Tester) AwaitEvent(events chan *event.Event) *event.Event {
	o.t.Helper()

	select {
	case e := <-events:
		return e
	case <-time.After(constant.FixtureEventTimeout):
		o.t.Fatal("no event arrived")

		return nil
	}
}

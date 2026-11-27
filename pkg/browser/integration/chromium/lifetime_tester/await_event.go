//go:build browser

package lifetime_tester

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium/event"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"time"
)

func (o *Tester) AwaitEvent(events chan *event.Event) *event.Event {
	o.t.Helper()

	select {
	case e := <-events:
		return e
	case <-time.After(constant.ChromiumFixtureEventTimeout):
		o.t.Fatal("no event arrived")

		return nil
	}
}

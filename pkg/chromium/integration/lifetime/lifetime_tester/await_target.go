//go:build browser

package lifetime_tester

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/event"
	"time"
)

func (o *Tester) AwaitTarget(
	events chan *event.Event,
	identifier string,
) {
	o.t.Helper()
	deadline := time.After(constant.FixtureEventTimeout)

	for {
		select {
		case e := <-events:
			if e.TargetIdentifier == identifier {
				return
			}
		case <-deadline:
			o.t.Fatalf("no event for target %s", identifier)
		}
	}
}

//go:build browser

package lifetime_tester

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium/event"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"time"
)

func (o *Tester) AwaitTarget(
	events chan *event.Event,
	identifier string,
) {
	o.t.Helper()
	deadline := time.After(constant.ChromiumFixtureEventTimeout)

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

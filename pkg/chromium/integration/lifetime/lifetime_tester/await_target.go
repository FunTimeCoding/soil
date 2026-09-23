//go:build browser

package lifetime_tester

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/event"
	"testing"
	"time"
)

func AwaitTarget(
	t *testing.T,
	events chan *event.Event,
	identifier string,
) {
	t.Helper()
	deadline := time.After(constant.FixtureEventTimeout)

	for {
		select {
		case e := <-events:
			if e.TargetIdentifier == identifier {
				return
			}
		case <-deadline:
			t.Fatalf("no event for target %s", identifier)
		}
	}
}

//go:build browser

package lifetime_tester

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/event"
	"testing"
	"time"
)

func AwaitEvent(
	t *testing.T,
	events chan *event.Event,
) *event.Event {
	t.Helper()

	select {
	case e := <-events:
		return e
	case <-time.After(constant.FixtureEventTimeout):
		t.Fatal("no event arrived")

		return nil
	}
}

//go:build browser

package lifetime

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"github.com/funtimecoding/soil/pkg/chromium/integration/lifetime/lifetime_tester"
	"testing"
)

func TestWatchRecordsTargetDestroyedByAnotherClient(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	destroyed := lifetime_tester.WatchKind(t, s, constant.EventKindDestroyed)
	other := chromium.New("localhost", s.Port)
	assert.FatalOnError(t, other.CloseTab(identifier))
	lifetime_tester.AwaitTarget(t, destroyed, identifier)
	other.Close()
}

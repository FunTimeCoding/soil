//go:build browser

package lifetime

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"github.com/funtimecoding/soil/pkg/chromium/integration/lifetime/lifetime_tester"
	"testing"
	"time"
)

func TestWatchRecordsTargetDestroyed(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	destroyed := lifetime_tester.WatchKind(t, s, constant.EventKindDestroyed)
	assert.FatalOnError(t, s.Client.CloseTab(identifier))
	lifetime_tester.AwaitTarget(t, destroyed, identifier)
}

func TestTargetCachePrunesOnDestroy(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	assert.StringContains(t, "quiet", s.Body(identifier))
	assert.Integer(t, 1, s.Client.TargetCount())
	assert.FatalOnError(t, s.Client.CloseTab(identifier))
	time.Sleep(2 * time.Second)
	assert.Integer(t, 0, s.Client.TargetCount())
}

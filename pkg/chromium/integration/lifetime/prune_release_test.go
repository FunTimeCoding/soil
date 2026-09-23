//go:build browser

package lifetime

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"github.com/funtimecoding/soil/pkg/chromium/integration/lifetime/lifetime_tester"
	"runtime"
	"testing"
	"time"
)

func TestPruneReleasesTargetGoroutines(t *testing.T) {
	s := base.New(t)
	lifetime_tester.Cycle(t, s)
	time.Sleep(constant.FixtureCloseSettlePeriod)
	before := runtime.NumGoroutine()

	for i := 0; i < constant.FixtureSharedCycleCount; i++ {
		lifetime_tester.Cycle(t, s)
	}

	time.Sleep(constant.FixtureCloseSettlePeriod)
	runtime.GC()
	after := runtime.NumGoroutine()
	assert.True(t, after-before < constant.FixtureSharedCycleCount)
	assert.Integer(t, 0, s.Client.TargetCount())
}

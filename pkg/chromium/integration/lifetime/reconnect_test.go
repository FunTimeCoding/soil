//go:build browser

package lifetime

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"github.com/funtimecoding/soil/pkg/errors"
	"runtime"
	"testing"
	"time"
)

func TestReconnectReleasesAbandonedTargetGoroutines(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	assert.StringContains(t, "quiet", s.Body(identifier))
	time.Sleep(constant.FixtureCloseSettlePeriod)
	before := runtime.NumGoroutine()

	for i := 0; i < constant.FixtureSharedCycleCount; i++ {
		errors.LogOnError(chromedp.Cancel(s.Client.Context()))
		assert.StringContains(t, "quiet", s.Body(identifier))
	}

	time.Sleep(constant.FixtureCloseSettlePeriod)
	runtime.GC()
	after := runtime.NumGoroutine()
	assert.True(t, after-before < constant.FixtureSharedCycleCount)
	assert.Integer(t, 1, s.Client.TargetCount())
	s.AssertTabAlive(identifier)
}

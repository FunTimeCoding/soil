//go:build browser

package lifetime

import (
	"context"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"github.com/funtimecoding/soil/pkg/chromium/integration/lifetime/lifetime_tester"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"runtime"
	"testing"
	"time"
)

func TestBodyKeepsTabOpen(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	assert.StringContains(t, "quiet", s.Body(identifier))
	s.AssertTabAlive(identifier)
}

func TestBodyOnStalledPageKeepsTabOpen(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureStalledRoute)
	assert.StringContains(t, "stalled", s.Body(identifier))
	s.AssertTabAlive(identifier)
}

func TestClientCloseAfterAcquireKeepsTabOpen(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	assert.StringContains(t, "quiet", s.Body(identifier))
	s.Client.Close()
	s.AssertTabAlive(identifier)
}

func TestClientCloseAfterAcquireOnBusyPageKeepsTabOpen(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureBusyRoute)
	assert.StringContains(t, "busy", s.Body(identifier))
	s.Client.Close()
	s.AssertTabAlive(identifier)
}

func TestClientCloseLeavesAcquiredContextUncancelled(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	x := s.Client.TargetContext(identifier)
	assert.StringContains(t, "quiet", s.Body(identifier))
	s.Client.Close()
	assert.Nil(t, x.Err())
	s.AssertTabAlive(identifier)
}

func TestClientCloseWithoutAcquireKeepsTabOpen(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	fresh := chromium.New("localhost", s.Port)
	fresh.Close()
	s.AssertTabAlive(identifier)
}

func TestDetachCarriesSessionOnly(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	detached := lifetime_tester.WatchKind(t, s, constant.EventKindDetached)
	x, cancel := chromedp.NewContext(
		s.Client.Context(),
		chromedp.WithTargetID(target.ID(identifier)),
	)
	var body string
	s.Client.RunContext(x, chromedp.OuterHTML(constant.BodySelector, &body))
	cancel()
	e := lifetime_tester.AwaitEvent(t, detached)
	assert.String(t, "", e.TargetIdentifier)
	assert.True(t, e.SessionIdentifier != "")
}

func TestWatchRecordsTargetDestroyedByAnotherClient(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	destroyed := lifetime_tester.WatchKind(t, s, constant.EventKindDestroyed)
	other := chromium.New("localhost", s.Port)
	assert.FatalOnError(t, other.CloseTab(identifier))
	lifetime_tester.AwaitTarget(t, destroyed, identifier)
	other.Close()
}

func TestProtocolReadsBorrowedTab(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	p := protocol.New(s.Client, constant.FixtureQuietRoute)
	assert.StringContains(t, "quiet", p.Body())
	assert.Integer(t, 1, s.Client.TargetCount())
	s.AssertTabAlive(identifier)
}

func TestProtocolAbsentTabLeavesClientUsable(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	assert.String(
		t,
		"tab not found",
		lifetime_tester.ProtocolAbsent(t, s.Client),
	)
	assert.Integer(t, 0, s.Client.TargetCount())
	assert.NotNil(t, s.Client.MustTabByHost(constant.FixtureQuietRoute))
	s.AssertTabAlive(identifier)
}

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

func TestSharedClientReuseKeepsTab(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)

	for i := 0; i < constant.FixtureSharedCycleCount; i++ {
		assert.StringContains(t, "quiet", s.Body(identifier))
	}

	assert.Integer(t, 1, s.Client.TargetCount())
	s.AssertTabAlive(identifier)
}

func TestCancelDerivedTargetContextKeepsTab(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	x, cancel := context.WithCancel(s.Client.TargetContext(identifier))
	var body string
	s.Client.RunContext(x, chromedp.OuterHTML(constant.BodySelector, &body))
	assert.StringContains(t, "quiet", body)
	cancel()
	s.AssertTabAlive(identifier)
}

func TestCancelAcquiredTargetContextTakesTab(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	x, cancel := chromedp.NewContext(
		s.Client.Context(),
		chromedp.WithTargetID(target.ID(identifier)),
	)
	var body string
	s.Client.RunContext(x, chromedp.OuterHTML(constant.BodySelector, &body))
	assert.StringContains(t, "quiet", body)
	cancel()
	s.AssertTabGone(identifier)
}

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

func TestFirstCallUnderDerivedDeadlineStrandsTarget(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	x, cancel := context.WithTimeout(
		s.Client.TargetContext(identifier),
		constant.FixtureEventTimeout,
	)
	var sum int
	assert.FatalOnError(
		t,
		chromedp.Run(x, chromedp.Evaluate(constant.FixtureSum, &sum)),
	)
	cancel()
	s.AssertTabAlive(identifier)
	p := protocol.NewIdentifier(
		s.Client,
		identifier,
	).WithTimeout(constant.FixtureCallTimeout)
	stranded := p.Evaluate(constant.FixtureSum, &sum)
	assert.True(t, timeout.Is(stranded))
}

func TestEvaluateTimeoutKeepsTabUsable(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	p := protocol.NewIdentifier(
		s.Client,
		identifier,
	).WithTimeout(constant.FixtureCallTimeout)
	var result any
	hung := p.EvaluatePromise(constant.FixtureHungPromise, &result)
	assert.True(t, timeout.Is(hung))
	assert.StringContains(t, "browser tab: did not answer within", hung.Error())
	s.AssertTabAlive(identifier)
	var sum int
	answered := p.Evaluate(constant.FixtureSum, &sum)
	assert.FatalOnError(t, answered)
	assert.Integer(t, 2, sum)
}

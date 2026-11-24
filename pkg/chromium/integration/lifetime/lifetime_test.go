//go:build browser

package lifetime

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"github.com/funtimecoding/soil/pkg/chromium/integration/lifetime/lifetime_tester"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
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
	x, f := s.Client.TargetContext(identifier)
	assert.FatalOnError(t, f)
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
	w := lifetime_tester.New(t)
	detached := w.WatchKind(s, constant.EventKindDetached)
	x, cancel := chromedp.NewContext(
		s.Client.Context(),
		chromedp.WithTargetID(target.ID(identifier)),
	)
	_, f := chromedp.Run(x, chromedp.OuterHTML(constant.BodySelector))
	assert.FatalOnError(t, f)
	cancel()
	e := w.AwaitEvent(detached)
	assert.String(t, "", e.TargetIdentifier)
	assert.True(t, e.SessionIdentifier != "")
}

func TestWatchRecordsTargetDestroyedByAnotherClient(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	w := lifetime_tester.New(t)
	destroyed := w.WatchKind(s, constant.EventKindDestroyed)
	other := chromium.New("localhost", s.Port)
	assert.FatalOnError(t, other.CloseTab(identifier))
	w.AwaitTarget(destroyed, identifier)
	other.Close()
}

func TestProtocolReadsBorrowedTab(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	p := protocol.New(s.Client, constant.FixtureQuietRoute)
	assert.StringContains(t, "quiet", p.MustBody())
	assert.Integer(t, 1, s.Client.TargetCount())
	s.AssertTabAlive(identifier)
}

func TestProtocolAbsentTabLeavesClientUsable(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	assert.String(
		t,
		"tab not found",
		lifetime_tester.New(t).ProtocolAbsent(s.Client),
	)
	assert.Integer(t, 0, s.Client.TargetCount())
	assert.NotNil(t, s.Client.MustHostTab(constant.FixtureQuietRoute))
	s.AssertTabAlive(identifier)
}

func TestPruneReleasesTargetGoroutines(t *testing.T) {
	s := base.New(t)
	w := lifetime_tester.New(t)
	w.Cycle(s)
	time.Sleep(constant.FixtureCloseSettlePeriod)
	before := runtime.NumGoroutine()

	for i := 0; i < constant.FixtureSharedCycleCount; i++ {
		w.Cycle(s)
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
	target, f := s.Client.TargetContext(identifier)
	assert.FatalOnError(t, f)
	x, cancel := context.WithCancel(target)
	body, e := chromedp.Run(x, chromedp.OuterHTML(constant.BodySelector))
	assert.FatalOnError(t, e)
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
	body, e := chromedp.Run(x, chromedp.OuterHTML(constant.BodySelector))
	assert.FatalOnError(t, e)
	assert.StringContains(t, "quiet", body)
	cancel()
	s.AssertTabGone(identifier)
}

func TestWatchRecordsTargetDestroyed(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	w := lifetime_tester.New(t)
	destroyed := w.WatchKind(s, constant.EventKindDestroyed)
	assert.FatalOnError(t, s.Client.CloseTab(identifier))
	w.AwaitTarget(destroyed, identifier)
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
	target, f := s.Client.TargetContext(identifier)
	assert.FatalOnError(t, f)
	x, cancel := context.WithTimeout(target, constant.FixtureEventTimeout)
	_, e := chromedp.Run(x, chromedp.Evaluate[int](constant.FixtureSum))
	assert.FatalOnError(t, e)
	cancel()
	s.AssertTabAlive(identifier)
	p := protocol.NewIdentifier(
		s.Client,
		identifier,
	).WithTimeout(constant.FixtureCallTimeout)
	var sum int
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
	assert.True(t, connection.Classify(hung) == nil)
	assert.StringContains(t, "browser tab: did not answer within", hung.Error())
	s.AssertTabAlive(identifier)
	var sum int
	answered := p.Evaluate(constant.FixtureSum, &sum)
	assert.FatalOnError(t, answered)
	assert.Integer(t, 2, sum)
}

func TestResumePageRunsFrozenTimers(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	x, e := s.Client.AcquireTarget(identifier)
	assert.FatalOnError(t, e)
	_, e = chromedp.Call(
		x,
		page.SetWebLifecycleState,
		page.SetWebLifecycleStateParams{
			State: page.SetWebLifecycleStateStateFrozen,
		},
	)
	assert.FatalOnError(t, e)
	p := protocol.NewIdentifier(
		s.Client,
		identifier,
	).WithTimeout(constant.FixtureCallTimeout)
	var stalled string
	assert.True(
		t,
		timeout.Is(p.EvaluatePromise(constant.FixtureTimer, &stalled)),
	)
	assert.FatalOnError(t, p.ResumePage())
	var fired string
	assert.FatalOnError(
		t,
		p.WithTimeout(constant.FixtureEventTimeout).EvaluatePromise(
			constant.FixtureTimer,
			&fired,
		),
	)
	assert.String(t, "timer", fired)
}

//go:build browser

package mechanic

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/browser_tester"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestEventAfterRequestResetsForm(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.NotifyPath))
	b.WaitVisible("#note-input")
	b.Fill("#note-input", "hello")
	b.Click("#receipt-post")
	b.WaitCondition(
		"document.querySelector('#receipt').textContent.indexOf('received hello') >= 0",
	)
	b.WaitCondition("document.querySelector('#note-input').value === ''")
}

func TestEventOutOfBandSwap(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.NotifyPath))
	b.WaitVisible("#note-input")
	b.Fill("#note-input", "hello")
	b.Click("#receipt-post")
	b.WaitCondition(
		"document.querySelector('#oob-summary').textContent.indexOf('sent 1') >= 0",
	)
	assert.StringContains(t, "sent 1", b.Text("#oob-summary"))
}

func TestEventFailureRaisesNotification(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.NotifyPath))
	b.WaitVisible("#fail-post")
	assert.Integer(t, 0, b.CountElements("#notifications .notification-error"))
	b.Click("#fail-post")
	b.WaitCondition(
		"document.querySelectorAll('#notifications > *').length > 0",
	)
	assert.String(t, "", b.Text("#receipt"))
}

func TestExtraSwapDelete(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.ExtraPath))
	b.WaitVisible("#remove-post")
	assert.Integer(t, 1, b.CountElements("#remove"))
	b.Click("#remove-post")
	b.WaitCondition("document.querySelector('#remove') === null")
}

func TestExtraSwapBeforeEnd(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.ExtraPath))
	b.WaitVisible("#append-post")
	assert.Integer(t, 0, b.CountElements("#append li"))
	b.Click("#append-post")
	b.WaitCondition("document.querySelectorAll('#append li').length === 1")
	b.Click("#append-post")
	b.WaitCondition("document.querySelectorAll('#append li').length === 2")
}

func TestExtraSwapNone(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.ExtraPath))
	b.WaitVisible("#quiet-post")
	b.Click("#quiet-post")
	b.WaitCondition("document.querySelector('#poll').textContent.length > 0")
	assert.String(t, "untouched", b.Text("#quiet"))
}

func TestExtraRequestHeaderBranch(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.ExtraPath))
	b.WaitCondition("document.querySelector('#branch').textContent.length > 0")
	assert.String(t, "fragment", b.Text("#branch"))
}

func TestExtraPolling(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.ExtraPath))
	b.WaitCondition(
		"document.querySelector('#poll').textContent.indexOf('poll 1') >= 0",
	)
	b.WaitCondition(
		"document.querySelector('#poll').textContent.indexOf('poll 3') >= 0",
	)
}

func TestExtraTriggerHeaderFiresEvent(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.ExtraPath))
	b.WaitVisible("#fire-post")
	assert.String(t, "", b.Text("#fire"))
	b.Click("#fire-post")
	b.WaitCondition("document.querySelector('#fire').textContent === 'heard'")
}

func TestExtraRedirectHeader(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.ExtraPath))
	b.WaitVisible("#redirect-post")
	b.Click("#redirect-post")
	b.WaitVisible("#pulse-post")
	assert.String(t, "pulse 0", b.Text("#pulse"))
}

func TestExtraRefreshHeader(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.ExtraPath))
	b.WaitVisible("#refresh-post")
	b.Evaluate("window.beforeRefresh = true", nil)
	b.Click("#refresh-post")
	b.WaitCondition("window.beforeRefresh === undefined")
}

func TestExtraIndicatorHiddenAtRest(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.ExtraPath))
	b.WaitVisible("#slow-post")
	assert.String(t, "0", b.Style(mark(constant.IndicatorMark), "opacity"))
}

func TestExtraIndicatorShowsWhileWaiting(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.ExtraPath))
	b.WaitVisible("#slow-post")
	b.Click("#slow-post")
	b.WaitCondition(
		"getComputedStyle(document.querySelector('#indicator-slot')).opacity === '1'",
	)
	b.WaitCondition(
		"document.querySelector('#slow-result').textContent === 'arrived'",
	)
	b.WaitCondition(
		"getComputedStyle(document.querySelector('#indicator-slot')).opacity === '0'",
	)
}

func TestExtraIndicatorSelfKeepsControlDimmed(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.ExtraPath))
	b.WaitVisible("#slow-post")
	b.Click("#slow-post")
	b.WaitCondition(
		"document.querySelector('#slow-post').classList.contains('htmx-request')",
	)
	b.WaitCondition(
		"document.querySelector('#slow-result').textContent === 'arrived'",
	)
	b.WaitCondition(
		"!document.querySelector('#slow-post').classList.contains('htmx-request')",
	)
}

func TestStreamConnects(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.StreamPath))
	b.WaitVisible("#pulse-post")
	assert.StringContains(t, "pulse 0", b.Text("#pulse"))
}

func TestStreamPushesNamedEvent(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.StreamPath))
	b.WaitVisible("#pulse-post")
	b.Click("#pulse-post")
	b.WaitCondition(
		"document.querySelector('#pulse').textContent.indexOf('pulse 1') >= 0",
	)
	assert.StringContains(t, "pulse 1", b.Text("#pulse"))
}

func TestStreamPushesSecondNamedEvent(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.StreamPath))
	b.WaitVisible("#pulse-post")
	b.Click("#pulse-post")
	assert.StringContains(t, "count 0", b.Text("#tick"))
	b.WaitCondition(
		"document.querySelector('#tick').textContent.indexOf('count 1') >= 0",
	)
	assert.StringContains(t, "count 1", b.Text("#tick"))
}

func TestSwapInner(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.RootPath))
	b.WaitVisible("#counter-post")
	assert.String(t, "0", b.Text("#counter"))
	b.Click("#counter-post")
	b.WaitCondition("document.querySelector('#counter').textContent === '1'")
	assert.String(t, "1", b.Text("#counter"))
}

func TestSwapGetReadsWithoutIncrement(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.RootPath))
	b.WaitVisible("#counter-get")
	b.Click("#counter-get")
	b.WaitCondition("document.querySelector('#counter') !== null")
	assert.String(t, "0", b.Text("#counter"))
}

func TestSwapOuterReplacesElement(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.RootPath))
	b.WaitVisible("#row-post")
	assert.String(t, "original", b.Text("#row"))
	b.Evaluate("window.confirm = function() { return true; }", nil)
	b.Click("#row-post")
	b.WaitCondition("document.querySelector('#row').textContent === 'replaced'")
	assert.Integer(t, 1, b.CountElements("#row"))
}

func TestSwapConfirmCancelBlocksRequest(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.RootPath))
	b.WaitVisible("#row-post")
	b.Evaluate("window.confirm = function() { return false; }", nil)
	b.Click("#row-post")
	b.WaitCondition("true")
	assert.String(t, "original", b.Text("#row"))
}

func TestTriggerLoad(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.TriggerPath))
	b.WaitCondition(
		"document.querySelector('#load-region').textContent.indexOf('term=load') >= 0",
	)
	assert.StringContains(t, "term=load", b.Text("#load-region"))
}

func TestTriggerLoadDelay(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.TriggerPath))
	assert.String(t, "", b.Text("#delayed"))
	b.WaitCondition(
		"document.querySelector('#delayed').textContent.indexOf('term=delayed') >= 0",
	)
}

func TestTriggerChange(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.TriggerPath))
	b.WaitVisible("#scope-select")
	b.Evaluate(
		"(function(){var s=document.querySelector('#scope-select');s.value='second';s.dispatchEvent(new Event('change'));})()",
		nil,
	)
	b.WaitCondition(
		"document.querySelector('#selection').textContent.indexOf('scope=second') >= 0",
	)
}

func TestTriggerKeyupIncludesAndVals(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.TriggerPath))
	b.WaitVisible("#term-input")
	b.Fill("#term-input", "abc")
	b.WaitCondition(
		"document.querySelector('#search').textContent.indexOf('term=abc') >= 0",
	)
	text := b.Text("#search")
	assert.StringContains(t, "term=abc", text)
	assert.StringContains(t, "scope=first", text)
	assert.StringContains(t, "origin=typed", text)
}

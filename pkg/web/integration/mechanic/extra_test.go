//go:build browser

package mechanic

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/browser_tester"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

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

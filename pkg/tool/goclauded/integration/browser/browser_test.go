//go:build browser

package browser

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/web/browser_tester"
	"testing"
	"time"
)

func TestAcquireTargetDoesNotCloseTab(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("about:blank")
	browser := chromedp.FromContext(b.Context).Browser
	identifier, e := target.CreateTarget("about:blank").Do(
		cdp.WithExecutor(b.Context, browser),
	)
	assert.FatalOnError(t, e)
	cached, _ := chromedp.NewContext(
		b.Context,
		chromedp.WithTargetID(identifier),
	)
	assert.FatalOnError(
		t,
		chromedp.Run(
			cached,
			chromedp.Navigate("data:text/html,<title>test</title>"),
		),
	)
	var title string
	assert.FatalOnError(t, chromedp.Run(cached, chromedp.Title(&title)))
	assert.String(t, "test", title)
	t.Run(
		"cached context survives repeated use",
		func(t *testing.T) {
			for i := 0; i < 5; i++ {
				done := make(chan error, 1)
				go func() {
					var v string
					done <- chromedp.Run(cached, chromedp.Title(&v))
				}()

				select {
				case f := <-done:
					assert.FatalOnError(t, f)
				case <-time.After(5 * time.Second):
					t.Fatalf("iteration %d: timed out", i)
				}
			}

			time.Sleep(time.Second)
			var alive string
			assert.FatalOnError(t, chromedp.Run(cached, chromedp.Title(&alive)))
			assert.String(t, "test", alive)
		},
	)
	t.Run(
		"context.WithTimeout cancel kills tab",
		func(t *testing.T) {
			fresh, _ := chromedp.NewContext(
				b.Context,
				chromedp.WithTargetID(target.ID(string(identifier))),
			)
			wrapped, cancel := context.WithTimeout(fresh, 10*time.Second)
			var v string
			assert.FatalOnError(t, chromedp.Run(wrapped, chromedp.Title(&v)))
			cancel()
			assert.String(t, "test", v)
			time.Sleep(time.Second)
			done := make(chan error, 1)
			go func() {
				var check string
				done <- chromedp.Run(fresh, chromedp.Title(&check))
			}()

			select {
			case f := <-done:
				assert.NotNil(t, f)
			case <-time.After(3 * time.Second):
			}
		},
	)
}

func TestBrowserSentinelDebug(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("http://localhost:8583/conversations")
	b.WaitReady(".sidebar-entry")
	var sentinelMarkup string
	b.Evaluate(
		"(() => { const el = document.querySelector('[hx-trigger=\"revealed\"]'); return el ? el.outerHTML : 'NOT FOUND'; })()",
		&sentinelMarkup,
	)
	console.Format("sentinel: %s\n", sentinelMarkup)
	var sidebarHeight float64
	b.Evaluate(
		"document.querySelector('.sidebar').scrollHeight",
		&sidebarHeight,
	)
	var sidebarClient float64
	b.Evaluate(
		"document.querySelector('.sidebar').clientHeight",
		&sidebarClient,
	)
	console.Format(
		"sidebar scrollHeight: %.0f, clientHeight: %.0f\n",
		sidebarHeight,
		sidebarClient,
	)
	b.ScrollToBottom(".sidebar")
	var scrollTop float64
	b.Evaluate("document.querySelector('.sidebar').scrollTop", &scrollTop)
	console.Format("after scroll, scrollTop: %.0f\n", scrollTop)
}

func TestBrowserSmoke(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("http://localhost:8583/conversations")
	var title string
	b.Evaluate("document.title", &title)
	console.Format("title: %q\n", title)
	var m string
	b.Evaluate("document.body.innerHTML.substring(0, 500)", &m)
	console.Format("body: %s\n", m)
}

func TestConversationsSidebarInfiniteScroll(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("http://localhost:8583/conversations")
	b.WaitReady(".sidebar-entry")
	initial := b.CountElements(".sidebar-entry")
	assert.True(t, initial > 0)
	console.Format("initial entries: %d\n", initial)
	b.ScrollToBottom(".sidebar")
	time.Sleep(2 * time.Second)
	after := b.CountElements(".sidebar-entry")
	console.Format("after scroll: %d\n", after)
	assert.True(t, after > initial)
}

func TestConversationsSidebarFilter(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("http://localhost:8583/conversations")
	b.WaitReady(".sidebar-entry")
	total := b.CountElements(".sidebar-entry")
	b.Evaluate(
		"document.querySelector('.sidebar-filter').value = 'goclauded'; document.querySelector('.sidebar-filter').dispatchEvent(new Event('input'))",
		nil,
	)
	time.Sleep(500 * time.Millisecond)
	var visible int
	b.Evaluate(
		"document.querySelectorAll('.sidebar-entry:not([style*=\"display: none\"])').length",
		&visible,
	)
	assert.True(t, visible > 0)
	assert.True(t, visible < total)
}

//go:build browser

package browser

import (
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
	created, e := chromedp.CallBrowser(
		b.Context,
		target.CreateTarget,
		target.CreateTargetParams{URL: "about:blank"},
	)
	assert.FatalOnError(t, e)
	cached, _ := chromedp.NewContext(
		b.Context,
		chromedp.WithTargetID(created.TargetID),
	)
	assert.FatalOnError(
		t,
		chromedp.Do(
			cached,
			chromedp.Navigate("data:text/html,<title>test</title>"),
		),
	)
	title, e := chromedp.Run(cached, chromedp.Title())
	assert.FatalOnError(t, e)
	assert.String(t, "test", title)
	t.Run(
		"cached context survives repeated use",
		func(t *testing.T) {
			for i := 0; i < 5; i++ {
				done := make(chan error, 1)
				go func() {
					_, f := chromedp.Run(cached, chromedp.Title())
					done <- f
				}()

				select {
				case f := <-done:
					assert.FatalOnError(t, f)
				case <-time.After(5 * time.Second):
					t.Fatalf("iteration %d: timed out", i)
				}
			}

			time.Sleep(time.Second)
			alive, f := chromedp.Run(cached, chromedp.Title())
			assert.FatalOnError(t, f)
			assert.String(t, "test", alive)
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

func TestConversationsSidebarSearch(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("https://localhost:8583/conversations")
	b.WaitReady(".sidebar-entry")
	b.Evaluate(
		"document.querySelector('.search-input').value = 'goclauded'; document.querySelector('.search-input').dispatchEvent(new Event('input', {bubbles: true}))",
		nil,
	)
	b.WaitCondition(
		"document.querySelectorAll('.search-result').length > 0 && document.querySelectorAll('.htmx-request, .htmx-swapping, .htmx-settling').length === 0",
	)
	assert.True(t, b.CountElements("#sidebar-entries mark") > 0)
	b.Evaluate("document.querySelector('.search-snippet').click()", nil)
	b.WaitCondition(
		"document.querySelectorAll('#panel .search-current').length == 1",
	)
	assert.True(t, b.CountElements("#panel .search-hit") > 0)
	b.Evaluate(
		"document.dispatchEvent(new KeyboardEvent('keydown', {key: 'n'}))",
		nil,
	)
	b.WaitCondition(
		"(document.getElementById('hit-counter').textContent || '').indexOf('hit ') === 0",
	)
}

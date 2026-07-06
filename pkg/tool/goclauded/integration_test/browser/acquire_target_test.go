//go:build browser

package browser

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/browser_tester"
	"testing"
	"time"
)

func TestAcquireTargetDoesNotCloseTab(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("about:blank")
	browser := chromedp.FromContext(b.Context).Browser
	identifier, e := target.CreateTarget(
		"about:blank",
	).Do(cdp.WithExecutor(b.Context, browser))
	assert.FatalOnError(t, e)
	// Cancel func discarded - calling it triggers chromedp cleanup that closes the tab
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
	assert.FatalOnError(
		t,
		chromedp.Run(cached, chromedp.Title(&title)),
	)
	assert.String(t, "test", title)
	t.Run(
		"cached context survives repeated use",
		func(t *testing.T) {
			// Goroutine + select: the timeout mechanism that doesn't touch the context chain
			for i := 0; i < 5; i++ {
				done := make(chan error, 1)
				go func() {
					var v string
					done <- chromedp.Run(
						cached,
						chromedp.Title(&v),
					)
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
			assert.FatalOnError(
				t,
				chromedp.Run(cached, chromedp.Title(&alive)),
			)
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
			wrapped, cancel := context.WithTimeout(
				fresh,
				10*time.Second,
			)
			var v string
			assert.FatalOnError(
				t,
				chromedp.Run(wrapped, chromedp.Title(&v)),
			)
			// cancel() propagates Done through the context chain; chromedp's
			// internal goroutine sees it and calls DetachFromTarget + CloseTarget
			cancel()
			assert.String(t, "test", v)
			time.Sleep(time.Second)
			done := make(chan error, 1)
			go func() {
				var check string
				done <- chromedp.Run(
					fresh,
					chromedp.Title(&check),
				)
			}()

			// Tab is either errored or unreachable - both confirm destruction
			select {
			case f := <-done:
				assert.NotNil(t, f)
			case <-time.After(3 * time.Second):
			}
		},
	)
}

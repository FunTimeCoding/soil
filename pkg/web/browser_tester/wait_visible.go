package browser_tester

import (
	"context"
	"github.com/chromedp/chromedp"
)

func (b *Browser) WaitVisible(selector string) {
	b.T.Helper()
	x, cancel := context.WithTimeout(b.Context, b.Timeout)
	defer cancel()

	if chromedp.Run(x, chromedp.WaitVisible(selector)) != nil {
		b.T.Fatalf("not visible within %s: %s", b.Timeout, selector)
	}
}

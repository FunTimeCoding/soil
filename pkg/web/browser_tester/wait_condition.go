package browser_tester

import (
	"context"
	"github.com/chromedp/chromedp"
)

func (b *Browser) WaitCondition(expression string) {
	b.T.Helper()
	var ready bool
	x, cancel := context.WithTimeout(b.Context, b.Timeout)
	defer cancel()

	if chromedp.Run(x, chromedp.Poll(expression, &ready)) != nil {
		b.T.Fatalf("condition not met within %s: %s", b.Timeout, expression)
	}
}

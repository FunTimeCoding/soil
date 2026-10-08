package browser_tester

import (
	"context"
	"github.com/chromedp/chromedp"
)

func (b *Browser) WaitCondition(expression string) {
	b.T.Helper()
	x, cancel := context.WithTimeout(b.Context, b.Timeout)
	defer cancel()

	if _, e := chromedp.Run(x, chromedp.Poll[bool](expression)); e != nil {
		b.T.Fatalf("condition not met within %s: %s", b.Timeout, expression)
	}
}

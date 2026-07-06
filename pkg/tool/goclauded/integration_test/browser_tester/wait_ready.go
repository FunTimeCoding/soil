package browser_tester

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (b *Browser) WaitReady(selector string) {
	b.T.Helper()
	errors.PanicOnError(
		chromedp.Run(b.Context, chromedp.WaitReady(selector)),
	)
}

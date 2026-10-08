package browser_tester

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (b *Browser) Fill(
	selector string,
	value string,
) {
	b.T.Helper()
	errors.PanicOnError(
		chromedp.Do(b.Context, chromedp.SendKeys(selector, value)),
	)
}

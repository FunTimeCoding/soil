package browser_tester

import (
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (b *Browser) ScrollToBottom(selector string) {
	b.T.Helper()
	errors.PanicOnError(
		chromedp.Do(
			b.Context,
			chromedp.Evaluate[chromedp.Void](
				fmt.Sprintf(
					"document.querySelector('%s').scrollTo(0, document.querySelector('%s').scrollHeight)",
					selector,
					selector,
				),
			),
		),
	)
}

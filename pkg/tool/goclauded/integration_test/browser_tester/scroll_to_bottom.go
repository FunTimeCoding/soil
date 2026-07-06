package browser_tester

import (
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (b *Browser) ScrollToBottom(selector string) {
	b.T.Helper()
	errors.PanicOnError(
		chromedp.Run(
			b.Context,
			chromedp.Evaluate(
				fmt.Sprintf(
					"document.querySelector('%s').scrollTo(0, document.querySelector('%s').scrollHeight)",
					selector,
					selector,
				),
				nil,
			),
		),
	)
}

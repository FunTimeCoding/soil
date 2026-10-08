package browser_tester

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (b *Browser) Evaluate(
	expression string,
	result any,
) {
	b.T.Helper()
	value, e := chromedp.Run(b.Context, chromedp.Evaluate[[]byte](expression))
	errors.PanicOnError(e)
	errors.PanicOnError(chromium.Decode(value, result))
}

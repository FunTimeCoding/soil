package protocol

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium"
)

func (p *Protocol) EvaluatePromise(
	expression string,
	result any,
) error {
	value, e := run(
		p,
		chromedp.Evaluate[[]byte](expression, chromedp.EvalAwaitPromise),
	)

	if e != nil {
		return e
	}

	return chromium.Decode(value, result)
}

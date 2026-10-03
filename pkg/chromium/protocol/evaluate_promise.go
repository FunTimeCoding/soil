package protocol

import (
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func (p *Protocol) EvaluatePromise(
	expression string,
	result any,
) error {
	return p.run(
		chromedp.Evaluate(
			expression,
			result,
			func(q *runtime.EvaluateParams) *runtime.EvaluateParams {
				return q.WithAwaitPromise(true)
			},
		),
	)
}

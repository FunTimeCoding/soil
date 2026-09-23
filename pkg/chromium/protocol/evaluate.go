package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) Evaluate(
	expression string,
	result any,
) error {
	return chromedp.Run(p.context, chromedp.Evaluate(expression, result))
}

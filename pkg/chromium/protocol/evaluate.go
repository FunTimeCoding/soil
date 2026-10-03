package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) Evaluate(
	expression string,
	result any,
) error {
	return p.run(chromedp.Evaluate(expression, result))
}

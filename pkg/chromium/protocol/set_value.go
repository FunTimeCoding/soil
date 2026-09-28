package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) SetValue(
	s string,
	value string,
) error {
	return chromedp.Run(
		p.context,
		chromedp.SetValue(s, value, chromedp.ByQuery),
	)
}

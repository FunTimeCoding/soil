package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) SetValue(
	s string,
	value string,
) error {
	return p.run(chromedp.SetValue(s, value, chromedp.ByQuery))
}

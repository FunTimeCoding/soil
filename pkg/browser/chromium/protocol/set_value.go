package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) SetValue(
	s string,
	value string,
) error {
	return p.do(chromedp.SetValue(chromedp.CSS(s), value))
}

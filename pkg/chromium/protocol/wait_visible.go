package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) WaitVisible(s string) error {
	return p.do(chromedp.WaitVisible(s))
}

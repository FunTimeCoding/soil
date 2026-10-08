package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) Outer(s string) (string, error) {
	return run(p, chromedp.OuterHTML(s))
}

package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) Location() (string, error) {
	return run(p, chromedp.Location())
}

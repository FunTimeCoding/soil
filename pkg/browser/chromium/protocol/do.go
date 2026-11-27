package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) do(steps ...chromedp.Action[chromedp.Void]) error {
	_, e := run(p, chromedp.Steps(steps...))

	return e
}

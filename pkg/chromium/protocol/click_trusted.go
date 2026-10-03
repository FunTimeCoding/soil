package protocol

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func (p *Protocol) ClickTrusted(s string) error {
	return p.run(
		chromedp.ActionFunc(
			func(v context.Context) error {
				return page.BringToFront().Do(v)
			},
		),
		chromedp.Click(s, chromedp.ByQuery),
	)
}

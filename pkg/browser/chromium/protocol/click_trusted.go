package protocol

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func (p *Protocol) ClickTrusted(s string) error {
	return p.do(
		chromedp.Func(
			func(
				v context.Context,
				t *chromedp.Target,
			) error {
				_, e := cdp.Call(v, t, page.BringToFront, cdp.Empty{})

				return e
			},
		),
		chromedp.Click(chromedp.CSS(s)),
	)
}

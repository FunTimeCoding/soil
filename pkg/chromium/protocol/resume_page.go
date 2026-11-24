package protocol

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func (p *Protocol) ResumePage() error {
	return p.do(
		chromedp.Func(
			func(
				v context.Context,
				t *chromedp.Target,
			) error {
				_, e := cdp.Call(
					v,
					t,
					page.SetWebLifecycleState,
					page.SetWebLifecycleStateParams{
						State: page.SetWebLifecycleStateStateActive,
					},
				)

				return e
			},
		),
	)
}

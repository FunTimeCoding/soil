package protocol

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
)

func (p *Protocol) Reload() error {
	return p.do(
		chromedp.Func(
			func(
				v context.Context,
				t *chromedp.Target,
			) error {
				_, e := cdp.Call(v, t, page.Reload, page.ReloadParams{})

				return e
			},
		),
		chromedp.WaitReady(constant.BodySelector),
	)
}

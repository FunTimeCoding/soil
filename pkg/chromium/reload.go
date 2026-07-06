package chromium

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Reload(x context.Context) {
	errors.PanicOnError(chromedp.Run(x, page.Reload()))
	errors.PanicOnError(
		chromedp.Run(
			x,
			chromedp.WaitReady(constant.BodySelector),
		),
	)
}

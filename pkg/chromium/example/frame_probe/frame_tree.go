package frame_probe

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/chromium"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func frameTree(
	c *chromium.Client,
	identifier string,
) *page.FrameTree {
	var result *page.FrameTree
	errors.PanicOnError(
		chromedp.Run(
			c.AcquireTarget(identifier),
			chromedp.ActionFunc(
				func(v context.Context) error {
					var e error
					result, e = page.GetFrameTree().Do(v)

					return e
				},
			),
		),
	)

	return result
}

package frame_probe

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/errors"
)

func frameTree(
	c *chromium.Client,
	identifier string,
) *page.FrameTree {
	result, e := chromedp.Call(
		c.AcquireTarget(identifier),
		page.GetFrameTree,
		cdp.Empty{},
	)
	errors.PanicOnError(e)

	return result.FrameTree
}

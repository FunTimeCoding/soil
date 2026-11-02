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
	x, e := c.AcquireTarget(identifier)
	errors.PanicOnError(e)
	result, f := chromedp.Call(x, page.GetFrameTree, cdp.Empty{})
	errors.PanicOnError(f)

	return result.FrameTree
}

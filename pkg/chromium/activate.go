package chromium

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func (c *Client) Activate(targetIdentifier string) {
	console.Line("  Activate")
	start := time.Now()
	console.Format("    Start %v\n", start.Format(constant.Micro))
	b, e := c.browser()
	errors.PanicOnError(e)
	t1 := time.Now()
	_, e = cdp.Call(
		c.context,
		b,
		target.ActivateTarget,
		target.ActivateTargetParams{TargetID: target.ID(targetIdentifier)},
	)
	errors.PanicOnError(e)
	console.Format("    ActivateTarget took %v\n", time.Since(t1))
	t2 := time.Now()
	_, e = chromedp.Call(
		c.TargetContext(targetIdentifier),
		page.Reload,
		page.ReloadParams{},
	)
	console.Format("    Reload took %v (error: %v)\n", time.Since(t2), e)
	errors.PanicOnError(e)
	console.Format("    Complete after: %v\n", time.Since(start))
}

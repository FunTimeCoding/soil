package chromium

import (
	"context"
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
	errors.PanicOnError(
		target.ActivateTarget(target.ID(targetIdentifier)).Do(
			cdp.WithExecutor(c.context, b),
		),
	)
	console.Format("    ActivateTarget took %v\n", time.Since(t1))
	t2 := time.Now()
	e = chromedp.Run(
		c.TargetContext(targetIdentifier),
		chromedp.ActionFunc(
			func(x context.Context) error {
				return page.Reload().Do(x)
			},
		),
	)
	console.Format("    Reload took %v (error: %v)\n", time.Since(t2), e)
	errors.PanicOnError(e)
	console.Format("    Complete after: %v\n", time.Since(start))
}

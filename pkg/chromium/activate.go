package chromium

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func (c *Client) Activate(targetIdentifier string) {
	fmt.Println("  Activate")
	start := time.Now()
	fmt.Printf("    Start %v\n", start.Format(library.Micro))
	b, e := c.browser()
	errors.PanicOnError(e)
	t1 := time.Now()
	errors.PanicOnError(
		target.ActivateTarget(target.ID(targetIdentifier)).Do(
			cdp.WithExecutor(c.context, b),
		),
	)
	fmt.Printf("    ActivateTarget took %v\n", time.Since(t1))
	t2 := time.Now()
	e = chromedp.Run(
		c.TargetContext(targetIdentifier),
		chromedp.ActionFunc(
			func(x context.Context) error {
				return page.Reload().Do(x)
			},
		),
	)
	fmt.Printf("    Reload took %v (error: %v)\n", time.Since(t2), e)
	errors.PanicOnError(e)
	fmt.Printf("    Complete after: %v\n", time.Since(start))
}

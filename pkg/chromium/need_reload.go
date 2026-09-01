package chromium

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (c *Client) NeedReload(
	identifier string,
	locator string,
) bool {
	c.reconnectIfNeeded()
	var found bool

	for _, g := range c.Targets() {
		if string(g.TargetID) == identifier {
			found = true

			break
		}
	}

	if !found {
		console.Line("  No target")

		return true
	}

	check, _ := chromedp.NewContext(
		c.context,
		chromedp.WithTargetID(target.ID(identifier)),
	)
	run, cancelRun := context.WithTimeout(check, 1*time.Second)
	defer cancelRun()

	if e := chromedp.Run(run); e != nil {
		if errors.Deadline(e) {
			console.Line("  Timeout run")

			return true
		}
	}

	resource, cancelResource := context.WithTimeout(check, 1*time.Second)
	defer cancelResource()

	if e := chromedp.Run(
		resource,
		chromedp.ActionFunc(
			func(o context.Context) error {
				t, e := page.GetResourceTree().Do(o)

				if e != nil {
					return e
				}

				_, e = page.GetResourceContent(t.Frame.ID, locator).Do(o)

				return e
			},
		),
	); e != nil {
		if errors.Deadline(e) {
			console.Line("  Timeout resource")

			return true
		}
	}

	return false
}

package protocol

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (p *Protocol) NeedReload(locator string) bool {
	run, cancelRun := context.WithTimeout(p.context, 1*time.Second)
	defer cancelRun()

	if e := chromedp.Run(run); e != nil {
		if errors.Deadline(e) {
			console.Line("  Timeout run")

			return true
		}
	}

	resource, cancelResource := context.WithTimeout(p.context, 1*time.Second)
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

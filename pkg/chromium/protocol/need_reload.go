package protocol

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (p *Protocol) NeedReload(locator string) bool {
	attach, cancelAttach := context.WithTimeout(p.context, 1*time.Second)
	defer cancelAttach()

	if e := chromedp.Do(attach); e != nil {
		if errors.Deadline(e) {
			console.Line("  Timeout run")

			return true
		}
	}

	resource, cancelResource := context.WithTimeout(p.context, 1*time.Second)
	defer cancelResource()

	if e := chromedp.Do(
		resource,
		chromedp.Func(
			func(
				o context.Context,
				s *chromedp.Target,
			) error {
				t, e := cdp.Call(o, s, page.GetResourceTree, cdp.Empty{})

				if e != nil {
					return e
				}

				_, e = cdp.Call(
					o,
					s,
					page.GetResourceContent,
					page.GetResourceContentParams{
						FrameID: t.FrameTree.Frame.ID,
						URL:     locator,
					},
				)

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

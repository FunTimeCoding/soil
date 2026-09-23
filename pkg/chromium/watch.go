package chromium

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/event"
)

func (c *Client) Watch(observe func(*event.Event)) error {
	chromedp.ListenBrowser(
		c.context,
		func(v any) {
			switch e := v.(type) {
			case *target.EventTargetCreated:
				observe(
					event.New(
						constant.EventKindCreated,
						string(e.TargetInfo.TargetID),
						e.TargetInfo.URL,
					),
				)
			case *target.EventTargetInfoChanged:
				observe(
					event.New(
						constant.EventKindChanged,
						string(e.TargetInfo.TargetID),
						e.TargetInfo.URL,
					),
				)
			case *target.EventTargetDestroyed:
				observe(
					event.New(
						constant.EventKindDestroyed,
						string(e.TargetID),
						"",
					),
				)
			case *target.EventAttachedToTarget:
				observe(
					event.New(
						constant.EventKindAttached,
						string(e.TargetInfo.TargetID),
						e.TargetInfo.URL,
					),
				)
			case *target.EventDetachedFromTarget:
				observe(
					event.NewSession(
						constant.EventKindDetached,
						string(e.SessionID),
					),
				)
			}
		},
	)
	b, e := c.browser()

	if e != nil {
		return e
	}

	return target.SetDiscoverTargets(true).Do(cdp.WithExecutor(c.context, b))
}

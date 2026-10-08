package chromium

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/event"
)

func (c *Client) Watch(observe func(*event.Event)) error {
	go forward(
		chromedp.BrowserEvents(c.context, target.TargetCreated),
		func(v target.EventTargetCreated) *event.Event {
			return event.New(
				constant.EventKindCreated,
				string(v.TargetInfo.TargetID),
				v.TargetInfo.URL,
			)
		},
		observe,
	)
	go forward(
		chromedp.BrowserEvents(c.context, target.TargetInfoChanged),
		func(v target.EventTargetInfoChanged) *event.Event {
			return event.New(
				constant.EventKindChanged,
				string(v.TargetInfo.TargetID),
				v.TargetInfo.URL,
			)
		},
		observe,
	)
	go forward(
		chromedp.BrowserEvents(c.context, target.TargetDestroyed),
		func(v target.EventTargetDestroyed) *event.Event {
			return event.New(
				constant.EventKindDestroyed,
				string(v.TargetID),
				"",
			)
		},
		observe,
	)
	go forward(
		chromedp.BrowserEvents(c.context, target.AttachedToTarget),
		func(v target.EventAttachedToTarget) *event.Event {
			return event.New(
				constant.EventKindAttached,
				string(v.TargetInfo.TargetID),
				v.TargetInfo.URL,
			)
		},
		observe,
	)
	go forward(
		chromedp.BrowserEvents(c.context, target.DetachedFromTarget),
		func(v target.EventDetachedFromTarget) *event.Event {
			return event.NewSession(
				constant.EventKindDetached,
				string(v.SessionID),
			)
		},
		observe,
	)
	b, e := c.browser()

	if e != nil {
		return e
	}

	_, e = cdp.Call(
		c.context,
		b,
		target.SetDiscoverTargets,
		target.SetDiscoverTargetsParams{Discover: true},
	)

	return e
}

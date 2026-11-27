package chromium

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/browser/chromium/event"
	"github.com/funtimecoding/soil/pkg/browser/constant"
)

func (c *Client) subscribe(observe func(*event.Event)) error {
	go forward(
		chromedp.BrowserEvents(c.context, target.TargetCreated),
		func(v target.EventTargetCreated) *event.Event {
			return event.New(
				constant.ChromiumEventKindCreated,
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
				constant.ChromiumEventKindChanged,
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
				constant.ChromiumEventKindDestroyed,
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
				constant.ChromiumEventKindAttached,
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
				constant.ChromiumEventKindDetached,
				string(v.SessionID),
			)
		},
		observe,
	)
	b, e := c.connectedBrowser()

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

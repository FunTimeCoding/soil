package chromium

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func (c *Client) ClickNode(
	x context.Context,
	backendNodeIdentifier int64,
) error {
	return chromedp.Run(
		x,
		chromedp.ActionFunc(
			func(v context.Context) error {
				o, e := dom.ResolveNode().WithBackendNodeID(
					cdp.BackendNodeID(backendNodeIdentifier),
				).Do(v)

				if e != nil {
					return e
				}

				_, exception, e := runtime.CallFunctionOn(
					"function() { this.click() }",
				).WithObjectID(o.ObjectID).Do(v)

				if e != nil {
					return e
				}

				if exception != nil {
					return fmt.Errorf("click failed: %s", exception.Text)
				}

				return nil
			},
		),
	)
}

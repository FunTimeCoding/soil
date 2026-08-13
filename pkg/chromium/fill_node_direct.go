package chromium

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func (c *Client) fillNodeDirect(
	x context.Context,
	backendNodeID int64,
	value string,
) error {
	return chromedp.Run(
		x,
		chromedp.ActionFunc(
			func(v context.Context) error {
				o, e := dom.ResolveNode().WithBackendNodeID(
					cdp.BackendNodeID(backendNodeID),
				).Do(v)

				if e != nil {
					return e
				}

				_, exception, e := runtime.CallFunctionOn(
					"function(v) { this.focus(); this.value = v; this.dispatchEvent(new Event('input', {bubbles: true})) }",
				).WithObjectID(
					o.ObjectID,
				).WithArguments(
					[]*runtime.CallArgument{
						{Value: []byte(fmt.Sprintf(`"%s"`, value))},
					},
				).Do(v)

				if e != nil {
					return e
				}

				if exception != nil {
					return fmt.Errorf("fill failed: %s", exception.Text)
				}

				return nil
			},
		),
	)
}

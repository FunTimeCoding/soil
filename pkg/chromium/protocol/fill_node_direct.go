package protocol

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func (p *Protocol) fillNodeDirect(
	backendNodeIdentifier int64,
	value string,
) error {
	return p.do(
		chromedp.Func(
			func(
				v context.Context,
				t *chromedp.Target,
			) error {
				o, e := cdp.Call(
					v,
					t,
					dom.ResolveNode,
					dom.ResolveNodeParams{
						BackendNodeID: cdp.BackendNodeID(backendNodeIdentifier),
					},
				)

				if e != nil {
					return e
				}

				r, e := cdp.Call(
					v,
					t,
					runtime.CallFunctionOn,
					runtime.CallFunctionOnParams{
						FunctionDeclaration: "function(v) { this.focus(); this.value = v; this.dispatchEvent(new Event('input', {bubbles: true})) }",
						ObjectID:            o.Object.ObjectID,
						Arguments: []*runtime.CallArgument{
							{Value: []byte(fmt.Sprintf(`"%s"`, value))},
						},
					},
				)

				if e != nil {
					return e
				}

				if r.ExceptionDetails != nil {
					return fmt.Errorf(
						"fill failed: %s",
						r.ExceptionDetails.Text,
					)
				}

				return nil
			},
		),
	)
}

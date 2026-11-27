package protocol

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func (p *Protocol) fillNodeInsertText(
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
						FunctionDeclaration: `function() {
						this.focus();
						if (this.tagName.toLowerCase() === 'select') {
							return 'select';
						}
						document.execCommand('selectAll');
						return 'text';
					}`,
						ObjectID: o.Object.ObjectID,
					},
				)

				if e != nil {
					return e
				}

				if r.ExceptionDetails != nil {
					return fmt.Errorf(
						"fill focus failed: %s",
						r.ExceptionDetails.Text,
					)
				}

				_, e = cdp.Call(
					v,
					t,
					input.InsertText,
					input.InsertTextParams{Text: value},
				)

				return e
			},
		),
	)
}

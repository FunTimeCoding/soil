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
	return chromedp.Run(
		p.context,
		chromedp.ActionFunc(
			func(v context.Context) error {
				o, e := dom.ResolveNode().WithBackendNodeID(
					cdp.BackendNodeID(backendNodeIdentifier),
				).Do(v)

				if e != nil {
					return e
				}

				_, exception, e := runtime.CallFunctionOn(
					`function() {
						this.focus();
						if (this.tagName.toLowerCase() === 'select') {
							return 'select';
						}
						document.execCommand('selectAll');
						return 'text';
					}`,
				).WithObjectID(o.ObjectID).Do(v)

				if e != nil {
					return e
				}

				if exception != nil {
					return fmt.Errorf("fill focus failed: %s", exception.Text)
				}

				return input.InsertText(value).Do(v)
			},
		),
	)
}

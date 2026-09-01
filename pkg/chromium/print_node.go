package chromium

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) PrintNode(
	o context.Context,
	selector string,
	attribute []string,
) {
	var node []*cdp.Node
	errors.PanicOnError(
		chromedp.Run(o, chromedp.Nodes(selector, &node, chromedp.ByQueryAll)),
	)
	console.Format("Selector: %s\n", selector)

	for i, n := range node {
		console.Format("Index: %d\n", i)

		for _, a := range attribute {
			console.Format("  %s: %s\n", a, n.AttributeValue(a))
		}
	}
}

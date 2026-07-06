package chromium

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) PrintNode(
	o context.Context,
	selector string,
	attribute []string,
) {
	var node []*cdp.Node
	errors.PanicOnError(
		chromedp.Run(
			o,
			chromedp.Nodes(selector, &node, chromedp.ByQueryAll),
		),
	)
	fmt.Printf("Selector: %s\n", selector)

	for i, n := range node {
		fmt.Printf("Index: %d\n", i)

		for _, a := range attribute {
			fmt.Printf("  %s: %s\n", a, n.AttributeValue(a))
		}
	}
}

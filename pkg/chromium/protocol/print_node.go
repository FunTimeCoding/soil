package protocol

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/console"
)

func (p *Protocol) PrintNode(
	s string,
	attribute []string,
) {
	var result []*cdp.Node
	p.client.RunContext(
		p.context,
		chromedp.Nodes(s, &result, chromedp.ByQueryAll),
	)
	console.Format("Selector: %s\n", s)

	for i, n := range result {
		console.Format("Index: %d\n", i)
		console.Format("  XPath: %s\n", n.FullXPath())

		for _, a := range attribute {
			console.Format("  %s: %s\n", a, n.AttributeValue(a))
		}
	}
}

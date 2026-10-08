package protocol

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/console"
)

func (p *Protocol) PrintNode(
	s string,
	attribute []string,
) error {
	nodes, e := run(p, chromedp.Nodes(chromedp.CSSAll(s)))

	if e != nil {
		return e
	}

	console.Format("Selector: %s\n", s)

	for i, n := range nodes {
		console.Format("Index: %d\n", i)
		console.Format("  XPath: %s\n", n.FullXPath())

		for _, a := range attribute {
			console.Format("  %s: %s\n", a, n.AttributeValue(a))
		}
	}

	return nil
}

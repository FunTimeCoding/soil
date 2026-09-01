package protocol

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/funtimecoding/soil/pkg/console"
)

func Print(
	n *cdp.Node,
	attribute []string,
) {
	console.Format("  XPath: %s\n", n.FullXPath())

	for _, a := range attribute {
		console.Format("  %s: %s\n", a, n.AttributeValue(a))
	}
}

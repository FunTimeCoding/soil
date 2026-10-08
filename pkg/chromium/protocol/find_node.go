package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) FindNode(
	s string,
	index int,
) (*chromedp.Node, bool, error) {
	nodes, e := run(p, chromedp.Nodes(chromedp.CSSAll(s)))

	if e != nil {
		return nil, false, e
	}

	if index < 0 || index >= len(nodes) {
		return nil, false, nil
	}

	return nodes[index], true, nil
}

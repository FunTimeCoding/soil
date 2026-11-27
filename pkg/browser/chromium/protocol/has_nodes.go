package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) HasNodes(s string) (bool, error) {
	nodes, e := run(p, chromedp.Nodes(s, chromedp.AtLeast(0)))

	if e != nil {
		return false, e
	}

	return len(nodes) > 0, nil
}

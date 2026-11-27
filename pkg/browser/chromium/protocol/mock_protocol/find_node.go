package mock_protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) FindNode(
	_ string,
	_ int,
) (*chromedp.Node, bool, error) {
	return nil, false, nil
}

package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) ClickQuery(s string) error {
	return p.do(chromedp.Click(chromedp.CSS(s), chromedp.NodeVisible))
}

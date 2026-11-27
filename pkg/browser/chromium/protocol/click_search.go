package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) ClickSearch(s string) error {
	return p.do(chromedp.Click(chromedp.Search(s), chromedp.NodeVisible))
}

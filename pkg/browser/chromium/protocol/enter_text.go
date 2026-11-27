package protocol

import "github.com/chromedp/chromedp"

func (p *Protocol) EnterText(
	s string,
	text string,
) error {
	return p.do(
		chromedp.Click(s),
		chromedp.SendKeys(s, text),
		chromedp.KeyEvent("\r"),
	)
}

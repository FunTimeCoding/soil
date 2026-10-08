package chromium

import "github.com/chromedp/chromedp"

func (c *Client) Run(steps ...chromedp.Action[chromedp.Void]) {
	c.reconnectIfNeeded()
	c.RunContext(c.context, steps...)
}

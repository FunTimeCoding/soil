package chromium

import "github.com/chromedp/chromedp"

func (c *Client) Run(steps ...chromedp.Action[chromedp.Void]) error {
	if e := c.ready(); e != nil {
		return e
	}

	return c.RunContext(c.context, steps...)
}

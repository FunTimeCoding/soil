package chromium

import "github.com/chromedp/chromedp"

func (c *Client) browser() (*chromedp.Browser, error) {
	if e := c.ready(); e != nil {
		return nil, e
	}

	return c.connectedBrowser()
}

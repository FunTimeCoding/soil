package chromium

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
)

func (c *Client) CreateTab(l string) (string, error) {
	b, e := c.browser()

	if e != nil {
		return "", e
	}

	identifier, e := target.CreateTarget(l).Do(cdp.WithExecutor(c.context, b))

	if e != nil {
		return "", e
	}

	return string(identifier), nil
}

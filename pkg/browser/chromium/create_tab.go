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

	result, e := cdp.Call(
		c.context,
		b,
		target.CreateTarget,
		target.CreateTargetParams{URL: l},
	)

	if e != nil {
		return "", e
	}

	return string(result.TargetID), nil
}

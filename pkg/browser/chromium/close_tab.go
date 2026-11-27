package chromium

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
)

func (c *Client) CloseTab(identifier string) error {
	b, e := c.browser()

	if e != nil {
		return e
	}

	_, e = cdp.Call(
		c.context,
		b,
		target.CloseTarget,
		target.CloseTargetParams{TargetID: target.ID(identifier)},
	)

	return e
}

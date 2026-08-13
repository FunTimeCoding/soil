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

	return target.CloseTarget(target.ID(identifier)).Do(
		cdp.WithExecutor(c.context, b),
	)
}

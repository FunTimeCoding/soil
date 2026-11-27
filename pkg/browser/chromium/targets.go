package chromium

import (
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func (c *Client) Targets() ([]*target.Info, error) {
	if e := c.ready(); e != nil {
		return nil, e
	}

	return chromedp.Targets(c.context)
}

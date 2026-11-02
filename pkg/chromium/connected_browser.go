package chromium

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
)

func (c *Client) connectedBrowser() (*chromedp.Browser, error) {
	x := chromedp.FromContext(c.context)

	if x == nil || x.Browser == nil {
		return nil, unreachable.Format(
			"browser %s:%d: no connection",
			c.host,
			c.port,
		)
	}

	return x.Browser, nil
}

package chromium

import "github.com/chromedp/chromedp"

func (c *Client) ready() error {
	c.connecting.Lock()
	defer c.connecting.Unlock()

	if c.listening && c.context.Err() == nil {
		return nil
	}

	if c.context.Err() != nil {
		c.renew()
	}

	if _, e := chromedp.Targets(c.context); e != nil {
		c.renew()

		return c.dialFailure(e)
	}

	if e := c.listenTargets(); e != nil {
		c.renew()

		return c.dialFailure(e)
	}

	c.listening = true

	return nil
}

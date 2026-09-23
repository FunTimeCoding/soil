package chromium

import (
	"context"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func (c *Client) AcquireTarget(identifier string) context.Context {
	c.reconnectIfNeeded()
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if x, okay := c.targets[identifier]; okay {
		if x.Err() == nil {
			return x
		}

		delete(c.targets, identifier)
	}

	x, _ := chromedp.NewContext(
		context.WithoutCancel(c.context),
		chromedp.WithTargetID(target.ID(identifier)),
	)
	c.targets[identifier] = x

	return x
}

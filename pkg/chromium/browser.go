package chromium

import (
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
)

func (c *Client) browser() (*chromedp.Browser, error) {
	if c.context.Err() != nil {
		if e := c.reconnect(); e != nil {
			return nil, e
		}
	}

	x := chromedp.FromContext(c.context)

	if x == nil {
		return nil, unreachable.Format("chromedp context is nil")
	}

	if x.Browser != nil {
		return x.Browser, nil
	}

	_, e := chromedp.Targets(c.context)

	if e != nil {
		return nil, fmt.Errorf("browser reconnect: %w", e)
	}

	if x.Browser == nil {
		return nil, unreachable.Format("browser is nil after reconnect")
	}

	return x.Browser, nil
}

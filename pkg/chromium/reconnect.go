package chromium

import (
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
	"log/slog"
)

func (c *Client) reconnect() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.context.Err() == nil {
		return nil
	}

	c.cancel()

	for _, x := range c.targets {
		go func() {
			if e := chromedp.Cancel(x); !errors.Deadline(e) {
				errors.LogOnError(e)
			}
		}()
	}

	clear(c.targets)
	fresh, cancel := chromedp.NewContext(c.allocator)

	if _, e := chromedp.Targets(fresh); e != nil {
		cancel()

		return fmt.Errorf("cdp reconnect: %w", e)
	}

	c.context = fresh
	c.cancel = cancel

	if e := c.listenTargets(); e != nil {
		return fmt.Errorf("cdp reconnect listen: %w", e)
	}

	slog.Info("reconnected CDP context")

	return nil
}

package chromium

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) renew() {
	c.cancel()
	c.mutex.Lock()

	for _, x := range c.targets {
		go func() {
			if e := chromedp.Cancel(x); !errors.Deadline(e) {
				errors.LogOnError(e)
			}
		}()
	}

	clear(c.targets)
	c.mutex.Unlock()
	c.context, c.cancel = chromedp.NewContext(c.allocator)
	c.listening = false
}

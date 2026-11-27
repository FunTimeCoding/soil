package chromium

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/browser/chromium/event"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) listenTargets() error {
	return c.subscribe(
		func(e *event.Event) {
			if e.Kind != constant.ChromiumEventKindDestroyed {
				return
			}

			c.mutex.Lock()
			x, okay := c.targets[e.TargetIdentifier]
			delete(c.targets, e.TargetIdentifier)
			c.mutex.Unlock()

			if !okay {
				return
			}

			go func() {
				errors.LogOnError(chromedp.Cancel(x))
			}()
		},
	)
}

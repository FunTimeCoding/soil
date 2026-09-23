package chromium

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/event"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) listenTargets() error {
	return c.Watch(
		func(e *event.Event) {
			if e.Kind != constant.EventKindDestroyed {
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

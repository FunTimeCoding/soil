package connector

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
	"time"
)

func (c *Client) streamLoop(
	x context.Context,
	name string,
	kinds []string,
	events chan<- *stream_event.Event,
) {
	wait := constant.StreamRetryInterval
	last := ""

	for {
		if x.Err() != nil {
			return
		}

		delivered := c.streamOnce(x, name, kinds, last, events, &last)

		if x.Err() != nil {
			return
		}

		if delivered {
			wait = constant.StreamRetryInterval
		}

		select {
		case <-x.Done():
			return
		case <-time.After(wait):
		}

		wait *= 2

		if wait > constant.StreamRetryMaximum {
			wait = constant.StreamRetryMaximum
		}
	}
}

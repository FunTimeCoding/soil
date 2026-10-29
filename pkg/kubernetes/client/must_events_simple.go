package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"
)

func (c *Client) MustEventsSimple(
	normal bool,
	warning bool,
) []*event.Event {
	result, e := c.EventsSimple(normal, warning)
	errors.PanicOnError(e)

	return result
}

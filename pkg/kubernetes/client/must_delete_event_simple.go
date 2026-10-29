package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"
)

func (c *Client) MustDeleteEventSimple(v *event.Event) {
	errors.PanicOnError(c.DeleteEventSimple(v))
}

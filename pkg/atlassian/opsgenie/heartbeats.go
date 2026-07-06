package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/heartbeat"
)

func (c *Client) Heartbeats() []heartbeat.Heartbeat {
	r, e := c.userClient.Heartbeat.List(c.context)
	errors.PanicOnError(e)

	return r.Heartbeats
}

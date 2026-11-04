package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/notify_call"
)

func (c *Client) Notify(
	callsign string,
	source string,
	body string,
	d constant.Delivery,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return c.Failure
	}

	if e := c.failures[callsign]; e != nil {
		return e
	}

	c.notes = append(c.notes, notify_call.New(callsign, source, body, d))

	return nil
}

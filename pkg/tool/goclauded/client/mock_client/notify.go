package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/notify_call"

func (c *Client) Notify(
	callsign string,
	source string,
	body string,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return c.Failure
	}

	c.notes = append(c.notes, notify_call.New(callsign, source, body))

	return nil
}

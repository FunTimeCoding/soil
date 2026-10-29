package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"

func (c *Client) DeleteEventSimple(_ *event.Event) error {
	return c.Failure
}

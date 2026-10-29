package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"

func (c *Client) EventsSimple(_ bool, _ bool) ([]*event.Event, error) {
	return nil, c.Failure
}

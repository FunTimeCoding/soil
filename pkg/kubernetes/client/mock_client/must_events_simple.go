package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"

func (c *Client) MustEventsSimple(_ bool, _ bool) []*event.Event {
	return nil
}

package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"

func (c *Client) MustDeleteEventSimple(_ *event.Event) {}

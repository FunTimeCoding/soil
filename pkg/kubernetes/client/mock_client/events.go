package mock_client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"
)

func (c *Client) Events(_ *filter.Filter, _ int, _ string) ([]*event.Event, error) {
	return nil, c.Failure
}

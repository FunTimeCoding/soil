package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/filter"

func (c *Client) DeleteEvent(_ *filter.Filter) error {
	return c.Failure
}

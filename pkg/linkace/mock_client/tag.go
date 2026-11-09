package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/tag"

func (c *Client) Tag(_ string) (*tag.Tag, error) {
	return nil, c.Failure
}

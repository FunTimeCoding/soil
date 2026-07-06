package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/tag"

func (c *Client) MustTags() []*tag.Tag {
	return c.tags
}

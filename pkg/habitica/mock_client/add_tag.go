package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/tag"

func (c *Client) AddTag(t *tag.Tag) {
	c.tags = append(c.tags, t)
}

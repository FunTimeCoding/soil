package habitica

import "github.com/funtimecoding/soil/pkg/habitica/tag"

func (c *Client) Tags() ([]*tag.Tag, error) {
	var result []*tag.Tag
	e := c.basic.Get("/tags", nil, &result)

	return result, e
}

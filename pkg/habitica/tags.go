package habitica

import "github.com/funtimecoding/soil/pkg/habitica/tag"

func (c *Client) Tags() ([]*tag.Tag, error) {
	var result []*tag.Tag

	return result, c.get("/tags", &result)
}

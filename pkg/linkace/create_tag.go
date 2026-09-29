package linkace

import (
	"github.com/funtimecoding/soil/pkg/linkace/response"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

func (c *Client) CreateTag(name string) (*tag.Tag, error) {
	var r response.Tag

	if e := c.basic.Post(
		"tags",
		map[string]any{"name": name, "visibility": 3},
		&r,
	); e != nil {
		return nil, e
	}

	return tag.New(r, c.host), nil
}

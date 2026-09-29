package linkace

import (
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/linkace/response"
)

func (c *Client) CreateList(
	name string,
	description string,
) (*list.List, error) {
	var r response.List

	if e := c.basic.Post(
		"lists",
		map[string]any{
			"name":        name,
			"description": description,
			"visibility":  3,
		},
		&r,
	); e != nil {
		return nil, e
	}

	return list.New(r, c.host), nil
}

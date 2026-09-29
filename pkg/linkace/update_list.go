package linkace

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/linkace/response"
)

func (c *Client) UpdateList(
	identifier int,
	patch map[string]any,
) (*list.List, error) {
	var l response.List

	if e := c.basic.Patch(
		fmt.Sprintf("lists/%d", identifier),
		patch,
		&l,
	); e != nil {
		return nil, e
	}

	return list.New(l, c.host), nil
}

package linkace

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/response"
)

func (c *Client) UpdateLink(
	identifier int,
	patch map[string]any,
) (*link.Link, error) {
	var l response.Link

	if e := c.basic.Patch(
		fmt.Sprintf("links/%d", identifier),
		patch,
		&l,
	); e != nil {
		return nil, e
	}

	return link.New(l, c.host), nil
}

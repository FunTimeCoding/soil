package linkace

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/linkace/response"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

func (c *Client) UpdateTag(
	identifier int,
	patch map[string]any,
) (*tag.Tag, error) {
	var t response.Tag

	if e := c.basic.Patch(
		fmt.Sprintf("tags/%d", identifier),
		patch,
		&t,
	); e != nil {
		return nil, e
	}

	return tag.New(t, c.host), nil
}

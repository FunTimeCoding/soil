package netbox

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
)

func (c *Client) TagByName(n string) (*tag.Tag, error) {
	result, e := c.Tags()

	if e != nil {
		return nil, e
	}

	for _, t := range result {
		if t.Name == n {
			return t, nil
		}
	}

	return nil, not_found.New(constant.TagKey, n)
}

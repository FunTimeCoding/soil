package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
)

func (c *Client) Tags() ([]*tag.Tag, error) {
	if len(c.cache.Tags) != 0 {
		return c.cache.Tags, nil
	}

	result, _, e := c.client.ExtrasAPI.ExtrasTagsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.Tags = tag.NewSlice(result.Results)

	return c.cache.Tags, nil
}

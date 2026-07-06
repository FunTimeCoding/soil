package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateTag(name string) (*tag.Tag, error) {
	q := netbox.NewTagRequest(name, slug(name))
	result, _, e := c.client.ExtrasAPI.ExtrasTagsCreate(
		c.context,
	).TagRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.Tags = nil

	return tag.New(result), nil
}

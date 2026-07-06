package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
)

func (c *Client) Prefixes() ([]*prefix.Prefix, error) {
	if len(c.cache.Prefixes) != 0 {
		return c.cache.Prefixes, nil
	}

	result, _, e := c.client.IpamAPI.IpamPrefixesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.Prefixes = prefix.NewSlice(result.Results)

	return c.cache.Prefixes, nil
}

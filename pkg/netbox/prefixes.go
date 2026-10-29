package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) Prefixes() ([]*prefix.Prefix, error) {
	if len(c.cache.Prefixes) != 0 {
		return c.cache.Prefixes, nil
	}

	result, _, e := c.client.IpamAPI.IpamPrefixesList(c.context).Limit(
		constant.PageLimit,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	c.cache.Prefixes = prefix.NewSlice(result.Results)

	return c.cache.Prefixes, nil
}

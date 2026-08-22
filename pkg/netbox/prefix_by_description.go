package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
)

func (c *Client) PrefixByDescription(d string) (*prefix.Prefix, error) {
	all, e := c.Prefixes()

	if e != nil {
		return nil, e
	}

	var result []*prefix.Prefix

	for _, p := range all {
		if p.Description == d {
			result = append(result, p)
		}
	}

	if len(result) == 0 {
		return nil, not_found.Format("no prefix with description: %s", d)
	}

	if len(result) > 1 {
		return nil, ambiguous.Format(
			"expected 1 prefix with description %s, got %d",
			d,
			len(result),
		)
	}

	return result[0], nil
}

package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
)

func (c *Client) PrefixByName(n string) (*prefix.Prefix, error) {
	all, e := c.Prefixes()

	if e != nil {
		return nil, e
	}

	var result []*prefix.Prefix

	for _, p := range all {
		if p.Name == n {
			result = append(result, p)
		}
	}

	if len(result) == 0 {
		return nil, not_found.New("prefix", n)
	}

	if len(result) > 1 {
		return nil, ambiguous.Format(
			"expected 1 prefix named %s, got %d",
			n,
			len(result),
		)
	}

	return result[0], nil
}

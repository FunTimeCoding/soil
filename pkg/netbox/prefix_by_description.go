package netbox

import (
	"fmt"
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

	if len(result) != 1 {
		return nil, fmt.Errorf(
			"expected 1 prefix with description %s, got %d",
			d,
			len(result),
		)
	}

	return result[0], nil
}

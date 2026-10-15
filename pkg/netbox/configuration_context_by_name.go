package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (c *Client) ConfigurationContextByName(
	n string,
) (*configuration_context.Context, error) {
	result, e := c.ConfigurationContexts()

	if e != nil {
		return nil, e
	}

	for _, v := range result {
		if v.Name == n {
			return v, nil
		}
	}

	return nil, not_found.New(constant.ConfigurationContextKind, n)
}

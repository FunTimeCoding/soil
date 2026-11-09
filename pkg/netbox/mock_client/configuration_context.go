package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (c *Client) ConfigurationContext(
	n string,
) (*configuration_context.Context, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, v := range c.contexts {
		if v.Name == n {
			return v, nil
		}
	}

	return nil, not_found.New(constant.ConfigurationContextKind, n)
}

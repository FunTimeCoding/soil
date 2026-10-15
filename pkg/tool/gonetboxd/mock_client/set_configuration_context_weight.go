package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
)

func (c *Client) SetConfigurationContextWeight(
	name string,
	weight int32,
) (*configuration_context.Context, error) {
	v, e := c.ConfigurationContextByName(name)

	if e != nil {
		return nil, e
	}

	v.Weight = weight
	c.record("SetConfigurationContextWeight", name, fmt.Sprint(weight))

	return v, nil
}

package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
)

func (c *Client) Services(
	_ context.Context,
	origin string,
) ([]client.Service, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	var result []client.Service

	for _, v := range c.ServiceList {
		if origin != "" && v.Origin != origin {
			continue
		}

		result = append(result, v)
	}

	return result, nil
}

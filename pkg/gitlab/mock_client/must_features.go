package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MustFeatures(x context.Context) []*gitlab.Feature {
	result, e := c.Features(x)
	errors.PanicOnError(e)

	return result
}

package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MustSearchBlob(
	x context.Context,
	query string,
) []*gitlab.Blob {
	result, e := c.SearchBlob(x, query)
	errors.PanicOnError(e)

	return result
}

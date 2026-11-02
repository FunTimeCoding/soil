package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustQuery(
	x context.Context,
	q string,
	response any,
) {
	errors.PanicOnError(c.Query(x, q, response))
}

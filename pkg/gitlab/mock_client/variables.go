package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/variable"
)

func (c *Client) Variables(
	_ context.Context,
	_ int64,
) ([]*variable.Variable, error) {
	return nil, nil
}

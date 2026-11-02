package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/variable"
)

func (c *Client) ProjectVariable(
	_ context.Context,
	_ int64,
	_ string,
) (*variable.Variable, error) {
	return nil, nil
}

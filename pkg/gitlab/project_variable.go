package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/variable"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) ProjectVariable(
	x context.Context,
	project int64,
	key string,
) (*variable.Variable, error) {
	result, _, e := c.client.ProjectVariables.GetVariable(
		project,
		key,
		nil,
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return variable.New(result), nil
}

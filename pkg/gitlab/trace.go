package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/system"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Trace(
	x context.Context,
	project int64,
	job int64,
) (string, error) {
	result, _, e := c.client.Jobs.GetTraceFile(
		project,
		job,
		gitlab.WithContext(x),
	)

	if e != nil {
		return "", wrapError(e)
	}

	return string(system.ReadAll(result)), nil
}

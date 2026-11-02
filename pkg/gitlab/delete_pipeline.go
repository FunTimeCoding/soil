package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) DeletePipeline(
	x context.Context,
	project int64,
	pipeline int64,
) error {
	r, e := c.client.Pipelines.DeletePipeline(
		project,
		pipeline,
		gitlab.WithContext(x),
	)

	if r != nil && r.StatusCode == 404 {
		return nil
	}

	return wrapError(e)
}

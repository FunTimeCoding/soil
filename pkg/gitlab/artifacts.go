package gitlab

import (
	"bytes"
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Artifacts(
	x context.Context,
	project int64,
	job int64,
) (*bytes.Reader, error) {
	result, _, e := c.client.Jobs.GetJobArtifacts(
		project,
		job,
		gitlab.WithContext(x),
	)

	return result, wrapError(e)
}

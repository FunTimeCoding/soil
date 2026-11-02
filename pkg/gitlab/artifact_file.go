package gitlab

import (
	"bytes"
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) ArtifactFile(
	x context.Context,
	project int64,
	job int64,
	path string,
) (*bytes.Reader, error) {
	result, _, e := c.client.Jobs.DownloadSingleArtifactsFile(
		project,
		job,
		path,
		gitlab.WithContext(x),
	)

	return result, wrapError(e)
}

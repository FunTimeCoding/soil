package gitlab

import (
	"github.com/funtimecoding/soil/pkg/system"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) ArtifactsFile(
	project int64,
	reference string,
	job string,
) (string, error) {
	// This only works for successful jobs
	reader, _, e := c.client.Jobs.DownloadArtifactsFile(
		project,
		reference,
		&gitlab.DownloadArtifactsFileOptions{Job: &job},
	)

	if e != nil {
		return "", e
	}

	return string(system.ReadAll(reader)), nil
}

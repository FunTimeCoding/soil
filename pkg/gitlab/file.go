package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) File(
	project int64,
	branch string,
	name string,
) (*gitlab.File, error) {
	result, r, e := c.client.RepositoryFiles.GetFile(
		project,
		name,
		&gitlab.GetFileOptions{Ref: &branch},
	)

	if r != nil && r.StatusCode == 404 {
		return nil, not_found.Format(
			"file not found: %s (branch %s, project %d)",
			name,
			branch,
			project,
		)
	}

	return result, wrapError(e)
}

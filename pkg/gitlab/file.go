package gitlab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
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
		return nil, fmt.Errorf(
			"file not found: %s (branch %s, project %d): %w",
			name,
			branch,
			project,
			constant.ErrorNotFound,
		)
	}

	return result, e
}

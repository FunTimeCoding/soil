package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/gitlab/file"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) File(
	project int64,
	branch string,
	name string,
) (*file.File, error) {
	o := &gitlab.GetFileOptions{}

	if branch != "" {
		o.Ref = &branch
	}

	result, r, e := c.client.RepositoryFiles.GetFile(project, name, o)

	if r != nil && r.StatusCode == 404 {
		return nil, not_found.Format(
			"file not found: %s (branch %s, project %d)",
			name,
			branch,
			project,
		)
	}

	if e != nil {
		return nil, wrapError(e)
	}

	return file.New(result), nil
}

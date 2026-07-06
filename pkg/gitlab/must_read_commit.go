package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustReadCommit(
	project int64,
	sha string,
) *gitlab.Commit {
	result, e := c.ReadCommit(project, sha)
	errors.PanicOnError(e)

	return result
}

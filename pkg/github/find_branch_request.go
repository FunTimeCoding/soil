package github

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/github/pull_request"
)

func (c *Client) FindBranchRequest(
	owner string,
	repository string,
	branch string,
) (*pull_request.Request, bool, error) {
	result, e := c.BranchRequest(owner, repository, branch)

	if e != nil {
		if not_found.Is(e) {
			return nil, false, nil
		}

		return nil, false, e
	}

	return result, true, nil
}

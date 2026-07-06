package github

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/pull_request"
)

func (c *Client) FindBranchRequest(
	owner string,
	repository string,
	branch string,
) (*pull_request.Request, bool, error) {
	result, e := c.BranchRequest(owner, repository, branch)

	if e != nil {
		if errors.Is(e, constant.ErrorNotFound) {
			return nil, false, nil
		}

		return nil, false, e
	}

	return result, true, nil
}

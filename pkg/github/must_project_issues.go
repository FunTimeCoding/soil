package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/google/go-github/v89/github"
)

func (c *Client) MustProjectIssues(
	owner string,
	name string,
) []*github.Issue {
	result, e := c.ProjectIssues(owner, name)
	errors.PanicOnError(e)

	return result
}

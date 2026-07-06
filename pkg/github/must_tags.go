package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/google/go-github/v88/github"
)

func (c *Client) MustTags(
	owner string,
	repository string,
) []*github.RepositoryTag {
	result, e := c.Tags(owner, repository)
	errors.PanicOnError(e)

	return result
}

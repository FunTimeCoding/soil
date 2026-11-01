package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/google/go-github/v93/github"
)

func (c *Client) MustTags(
	owner string,
	repository string,
) []*github.RepositoryTag {
	result, e := c.Tags(owner, repository)
	errors.PanicOnError(e)

	return result
}

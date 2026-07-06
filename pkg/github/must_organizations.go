package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/google/go-github/v88/github"
)

func (c *Client) MustOrganizations(user string) []*github.Organization {
	result, e := c.Organizations(user)
	errors.PanicOnError(e)

	return result
}

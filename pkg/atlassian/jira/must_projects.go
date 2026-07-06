package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustProjects() *jira.ProjectList {
	result, e := c.Projects()
	errors.PanicOnError(e)

	return result
}

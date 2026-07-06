package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustFillComments(v []*issue.Issue) {
	errors.PanicOnError(c.FillComments(v))
}

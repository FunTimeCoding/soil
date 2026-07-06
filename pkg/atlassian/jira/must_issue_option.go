package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/option"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustIssueOption() *option.Issue {
	result, e := c.IssueOption()
	errors.PanicOnError(e)

	return result
}

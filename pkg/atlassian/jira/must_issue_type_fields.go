package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustIssueTypeFields(t *jira.MetaIssueType) map[string]string {
	result, e := c.IssueTypeFields(t)
	errors.PanicOnError(e)

	return result
}

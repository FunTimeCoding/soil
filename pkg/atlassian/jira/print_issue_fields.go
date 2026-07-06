package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) PrintIssueFields(
	projectKey string,
	issueType string,
) {
	p, e := c.MetaProject(projectKey)
	errors.PanicOnError(e)
	t, f := c.MetaIssueType(p, issueType)
	errors.PanicOnError(f)
	fields, g := c.IssueTypeFields(t)
	errors.PanicOnError(g)

	for k, v := range fields {
		fmt.Printf("Field: %s = %s\n", k, v)
	}
}

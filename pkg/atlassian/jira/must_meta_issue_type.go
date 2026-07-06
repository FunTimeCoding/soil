package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustMetaIssueType(
	p *jira.MetaProject,
	issueType string,
) *jira.MetaIssueType {
	result, e := c.MetaIssueType(p, issueType)
	errors.PanicOnError(e)

	return result
}

package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustNewIssueUnverified(
	projectKey string,
	issueType string,
	summary string,
	description string,
) *jira.Issue {
	result, e := c.NewIssueUnverified(
		projectKey,
		issueType,
		summary,
		description,
	)
	errors.PanicOnError(e)

	return result
}

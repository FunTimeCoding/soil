package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func (c *Client) SearchLimit(
	limit int,
	query string,
	a ...any,
) ([]*issue.Issue, error) {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	page, _, e := c.client.Issue.SearchV2JQL(
		query,
		&jira.SearchOptionsV2{Fields: []string{"*all"}, MaxResults: limit},
	)

	if e != nil {
		return nil, e
	}

	o, f := c.IssueOption()

	if f != nil {
		return nil, f
	}

	result := issue.NewSlice(page, o)

	return c.enrichMany(result), nil
}

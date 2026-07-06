package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func (c *Client) Search(
	query string,
	a ...any,
) ([]*issue.Issue, error) {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	o, f := c.IssueOption()

	if f != nil {
		return nil, f
	}

	var result []*issue.Issue
	var token string

	for {
		page, r, e := c.client.Issue.SearchV2JQL(
			query,
			&jira.SearchOptionsV2{
				NextPageToken: token,
				Fields:        []string{"*all"},
				MaxResults:    constant.SearchLimit,
				Expand:        constant.ChangelogExpand,
			},
		)

		if e != nil {
			return nil, e
		}

		token = r.NextPageToken
		result = append(result, issue.NewSlice(page, o)...)

		if r.IsLast {
			break
		}
	}

	return c.enrichMany(result), nil
}

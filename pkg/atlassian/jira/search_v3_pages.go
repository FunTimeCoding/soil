package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/response"
)

func (c *Client) searchV3Pages(
	query string,
	fields string,
	expand string,
) ([]*response.Issue, error) {
	var result []*response.Issue
	var token string

	for {
		page, e := c.searchV3Page(
			constant.JiraBasicSearchLimit,
			token,
			query,
			fields,
			expand,
		)

		if e != nil {
			return nil, e
		}

		result = append(result, page.Issues...)

		if page.NextPageToken == "" {
			break
		}

		if page.NextPageToken == token {
			return nil, fmt.Errorf(
				"pagination token did not advance: %s",
				token,
			)
		}

		token = page.NextPageToken
	}

	return result, nil
}

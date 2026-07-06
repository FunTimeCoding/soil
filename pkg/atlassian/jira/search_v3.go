package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/request"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) SearchV3(
	query string,
	a ...any,
) ([]*response.Issue, error) {
	// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-search
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	var result []*response.Issue

	if true {
		var token = ""

		for {
			page, e := c.searchV3Page(constant.BasicSearchLimit, token, query)

			if e != nil {
				return nil, e
			}

			result = append(result, page.Issues...)

			if page.IsLast {
				break
			}

			token = page.NextPageToken
		}
	}

	if false {
		// Response: 500 {"message":"Cannot invoke \"java.util.List.size()\" because \"reconcileIssues\" is null","status-code":500,"stack-trace":""}
		status, r, e := c.basic.PostPath(
			"/rest/api/3/search/jql",
			notation.Encode(
				request.Search{
					MaxResults: constant.BasicSearchLimit,
					Jql:        query,
				},
				false,
			),
		)

		if e != nil {
			return nil, e
		}

		fmt.Printf("Response: %d %s\n", status, r)
	}

	// Do not enrich, otherwise watchedIssueKeys will be recursive.
	return result, nil
}

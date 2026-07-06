package jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (c *Client) SearchFull(
	query string,
	a ...any,
) ([]*issue.Issue, error) {
	// andygrunwald/go-jira does not yet support new search API
	// https://developer.atlassian.com/changelog/#CHANGE-2046
	// https://github.com/andygrunwald/go-jira/issues/715
	var result []*issue.Issue
	pages, f := c.SearchV3(query, a...)

	if f != nil {
		return nil, f
	}

	for _, i := range pages {
		v, e := c.Issue(i.Key)

		if e != nil {
			return nil, e
		}

		result = append(result, v)
	}

	return c.enrichMany(result), nil
}

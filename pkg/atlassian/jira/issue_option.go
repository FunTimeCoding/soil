package jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue/option"

func (c *Client) IssueOption() (*option.Issue, error) {
	if c.issueOption != nil {
		return c.issueOption, nil
	}

	m, e := c.FieldMap()

	if e != nil {
		return nil, e
	}

	var keys []string

	if c.watchedIssues {
		k, f := c.WatchedIssueKeys()

		if f != nil {
			return nil, f
		}

		keys = k
	}

	c.issueOption = option.New(c.locator, c.user, keys, c.closedStatus, m)
	c.issueOption.Verbose = c.verbose
	c.issueOption.WatchedLoaded = c.watchedIssues

	return c.issueOption, nil
}

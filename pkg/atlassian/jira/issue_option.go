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

	keys, f := c.WatchedIssueKeys()

	if f != nil {
		return nil, f
	}

	c.issueOption = option.New(
		c.locator,
		c.user,
		keys,
		c.closedStatus,
		m,
	)
	c.issueOption.Verbose = c.verbose

	return c.issueOption, nil
}

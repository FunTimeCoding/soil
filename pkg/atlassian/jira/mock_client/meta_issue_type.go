package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) MetaIssueType(
	*jira.MetaProject,
	string,
) (*jira.MetaIssueType, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return nil, nil
}

package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) MetaIssueType(
	*jira.MetaProject,
	string,
) (*jira.MetaIssueType, error) {
	return nil, nil
}

package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) MetaProject(string) (*jira.MetaProject, error) { return nil, nil }

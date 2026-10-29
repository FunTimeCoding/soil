package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) User() (*jira.User, error) { return nil, nil }

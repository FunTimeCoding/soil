package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) AddLink(_ *jira.IssueLink) error { return nil }

package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) UpdateNative(_ *jira.Issue) error { return nil }

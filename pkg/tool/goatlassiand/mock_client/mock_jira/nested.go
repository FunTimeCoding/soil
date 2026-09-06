package mock_jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Nested() *jira.Client { return nil }

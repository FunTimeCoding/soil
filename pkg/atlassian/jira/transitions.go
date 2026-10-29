package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Transitions(key string) ([]jira.Transition, error) {
	result, r, e := c.client.Issue.GetTransitionsWithContext(c.context, key)

	return result, WrapError(r, e)
}

package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Transitions(key string) ([]jira.Transition, error) {
	result, _, e := c.client.Issue.GetTransitionsWithContext(c.context, key)

	return result, wrapError(e)
}

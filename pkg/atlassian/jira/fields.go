package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Fields() ([]jira.Field, error) {
	result, _, e := c.client.Field.GetList()

	return result, wrapError(e)
}

package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Fields() ([]jira.Field, error) {
	result, r, e := c.client.Field.GetList()

	return result, WrapError(r, e)
}

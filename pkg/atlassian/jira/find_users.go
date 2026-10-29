package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) FindUsers(query string) ([]jira.User, error) {
	users, r, e := c.client.User.FindWithContext(c.context, query)

	if e != nil {
		return nil, WrapError(r, e)
	}

	return users, nil
}

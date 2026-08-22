package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) FindUsers(query string) ([]jira.User, error) {
	users, _, e := c.client.User.FindWithContext(c.context, query)

	if e != nil {
		return nil, wrapError(e)
	}

	return users, nil
}

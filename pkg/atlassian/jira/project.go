package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Project(key string) (*jira.Project, error) {
	result, r, e := c.client.Project.Get(key)

	return result, WrapError(r, e)
}

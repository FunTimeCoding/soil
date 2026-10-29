package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Projects() (*jira.ProjectList, error) {
	result, r, e := c.client.Project.GetList()

	return result, WrapError(r, e)
}

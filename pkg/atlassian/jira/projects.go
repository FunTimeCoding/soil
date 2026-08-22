package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Projects() (*jira.ProjectList, error) {
	result, _, e := c.client.Project.GetList()

	return result, wrapError(e)
}

package jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/basic"

func (c *Client) Basic() *basic.Client {
	return c.basic
}

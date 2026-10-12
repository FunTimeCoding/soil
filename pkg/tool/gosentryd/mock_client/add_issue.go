package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) AddIssue(
	identifier string,
	project string,
	status string,
) {
	i := response.NewIssue()
	i.Identifier = identifier
	i.Project.Identifier = project
	i.Status = status
	c.issues = append(c.issues, i)
}

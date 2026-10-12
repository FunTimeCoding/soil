package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) AddProject(
	identifier string,
	slug string,
) {
	p := response.NewProject()
	p.Identifier = identifier
	p.Slug = slug
	c.projects = append(c.projects, *p)
}

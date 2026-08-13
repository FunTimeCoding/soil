package jira

import "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"

func (c *Client) RequestTypes(
	desk int,
	group int,
) (*models.ProjectRequestTypePageScheme, error) {
	result, _, e := c.service.Request.Type.Gets(c.context, desk, group, 0, 50)

	return result, e
}

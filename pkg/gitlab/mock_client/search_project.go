package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"strings"
)

func (c *Client) SearchProject(
	_ context.Context,
	query string,
) ([]*project.Project, error) {
	var result []*project.Project

	for _, p := range c.projects {
		if strings.Contains(p.CombinedName(), query) {
			result = append(result, p)
		}
	}

	return result, nil
}

package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"strings"
)

func (c *Client) SearchProject(query string) ([]*project.Project, error) {
	var result []*project.Project

	for _, p := range c.projects {
		if strings.Contains(p.CombinedName(), query) {
			result = append(result, p)
		}
	}

	return result, nil
}

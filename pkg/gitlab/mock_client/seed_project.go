package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) SeedProject(
	identifier int64,
	namespace string,
	name string,
) {
	c.projects = append(
		c.projects,
		project.New(
			&gitlab.Project{
				ID:        identifier,
				Path:      name,
				Namespace: &gitlab.ProjectNamespace{Path: namespace},
			},
		),
	)
}

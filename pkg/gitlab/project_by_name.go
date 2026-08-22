package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) ProjectByName(
	namespace string,
	name string,
) (*project.Project, error) {
	result, _, e := c.client.Projects.ListProjects(
		&gitlab.ListProjectsOptions{Search: &name},
	)

	if e != nil {
		return nil, wrapError(e)
	}

	count := len(result)

	if count == 0 {
		return nil, not_found.Format(
			"project not found: %s/%s",
			namespace,
			name,
		)
	}

	if count == 1 {
		if result[0].Namespace.Path == namespace {
			return project.New(result[0]), nil
		}

		return nil, not_found.Format(
			"project not found: %s/%s",
			namespace,
			name,
		)
	}

	for _, l := range result {
		if l.Namespace.Path == namespace && l.Name == name {
			return project.New(l), nil
		}
	}

	return nil, not_found.Format(
		"no exact match for project %s/%s among %d results",
		namespace,
		name,
		count,
	)
}

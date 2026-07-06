package gitlab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
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
		return nil, e
	}

	count := len(result)

	if count == 0 {
		return nil, fmt.Errorf(
			"project not found: %s/%s: %w",
			namespace,
			name,
			constant.ErrorNotFound,
		)
	}

	if count == 1 {
		if result[0].Namespace.Path == namespace {
			return project.New(result[0]), nil
		}

		return nil, fmt.Errorf(
			"project not found: %s/%s: %w",
			namespace,
			name,
			constant.ErrorNotFound,
		)
	}

	for _, l := range result {
		if l.Namespace.Path == namespace && l.Name == name {
			return project.New(l), nil
		}
	}

	return nil, fmt.Errorf(
		"no exact match for project %s/%s among %d results: %w",
		namespace,
		name,
		count,
		constant.ErrorNotFound,
	)
}

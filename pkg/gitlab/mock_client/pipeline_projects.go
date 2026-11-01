package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (c *Client) PipelineProjects() ([]*project.Project, error) {
	return nil, c.failure
}

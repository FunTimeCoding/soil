package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (c *Client) MustProjects() []*project.Project {
	return nil
}

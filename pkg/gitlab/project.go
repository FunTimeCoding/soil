package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (c *Client) Project(identifier int64) (*project.Project, error) {
	if p, okay := c.projectCache[identifier]; okay {
		return p, nil
	}

	result, _, e := c.client.Projects.GetProject(identifier, nil)

	if e != nil {
		return nil, e
	}

	c.projectCache[identifier] = project.New(result)

	return c.projectCache[identifier], nil
}

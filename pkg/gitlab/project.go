package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (c *Client) Project(identifier int64) (*project.Project, error) {
	c.mutex.Lock()
	p, okay := c.projectCache[identifier]
	c.mutex.Unlock()

	if okay {
		return p, nil
	}

	result, _, e := c.client.Projects.GetProject(identifier, nil)

	if e != nil {
		return nil, wrapError(e)
	}

	p = project.New(result)
	c.mutex.Lock()
	c.projectCache[identifier] = p
	c.mutex.Unlock()

	return p, nil
}

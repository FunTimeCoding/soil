package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Project(
	x context.Context,
	identifier int64,
) (*project.Project, error) {
	c.mutex.Lock()
	p, okay := c.projectCache[identifier]
	c.mutex.Unlock()

	if okay {
		return p, nil
	}

	result, _, e := c.client.Projects.GetProject(
		identifier,
		nil,
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	p = project.New(result)
	c.mutex.Lock()
	c.projectCache[identifier] = p
	c.mutex.Unlock()

	return p, nil
}

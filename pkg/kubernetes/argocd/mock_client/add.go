package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/argocd/application"

func (c *Client) Add(a ...*application.Application) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.applications = append(c.applications, a...)
}

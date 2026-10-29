package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/argocd/application"

func (c *Client) Applications() ([]*application.Application, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return append([]*application.Application{}, c.applications...), nil
}

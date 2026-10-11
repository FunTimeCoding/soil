package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/argocd/application"

func (c *Client) Applications() ([]*application.Application, error) {
	return nil, nil
}

package mock_client

import "github.com/funtimecoding/soil/pkg/argocd/application"

func (c *Client) Applications() ([]*application.Application, error) {
	return nil, nil
}

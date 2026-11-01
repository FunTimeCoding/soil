package mock_client

import "github.com/funtimecoding/soil/pkg/github/repository"

func (c *Client) ActionRepository() ([]*repository.Repository, error) {
	return nil, c.Failure
}

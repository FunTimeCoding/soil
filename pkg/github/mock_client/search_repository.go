package mock_client

import "github.com/funtimecoding/soil/pkg/github/repository"

func (c *Client) SearchRepository(_ string, _ ...any) ([]*repository.Repository, error) {
	return nil, c.Failure
}

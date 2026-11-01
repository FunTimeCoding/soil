package mock_client

import "github.com/google/go-github/v93/github"

func (c *Client) Repositories(_ string) ([]*github.Repository, error) {
	return nil, c.Failure
}

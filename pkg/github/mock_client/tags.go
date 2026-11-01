package mock_client

import "github.com/google/go-github/v93/github"

func (c *Client) Tags(_ string, _ string) ([]*github.RepositoryTag, error) {
	return nil, c.Failure
}

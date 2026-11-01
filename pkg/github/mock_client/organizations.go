package mock_client

import "github.com/google/go-github/v93/github"

func (c *Client) Organizations(_ string) ([]*github.Organization, error) {
	return nil, c.Failure
}

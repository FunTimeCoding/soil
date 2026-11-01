package mock_client

import "github.com/google/go-github/v93/github"

func (c *Client) PullRequests(_ string, _ string) ([]*github.PullRequest, error) {
	return nil, c.Failure
}

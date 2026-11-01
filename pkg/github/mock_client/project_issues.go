package mock_client

import "github.com/google/go-github/v93/github"

func (c *Client) ProjectIssues(_ string, _ string) ([]*github.Issue, error) {
	return nil, c.Failure
}

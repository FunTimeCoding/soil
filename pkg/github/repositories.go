package github

import "github.com/google/go-github/v89/github"

func (c *Client) Repositories(owner string) ([]*github.Repository, error) {
	result, _, e := c.client.Repositories.ListByUser(
		c.context,
		owner,
		nil,
	)

	return result, e
}

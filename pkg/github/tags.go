package github

import "github.com/google/go-github/v89/github"

func (c *Client) Tags(
	owner string,
	repository string,
) ([]*github.RepositoryTag, error) {
	result, _, e := c.client.Repositories.ListTags(
		c.context,
		owner,
		repository,
		&github.ListOptions{PerPage: 1000},
	)

	return result, e
}

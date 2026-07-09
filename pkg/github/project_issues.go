package github

import "github.com/google/go-github/v89/github"

func (c *Client) ProjectIssues(
	owner string,
	name string,
) ([]*github.Issue, error) {
	result, _, e := c.client.Issues.ListByRepo(
		c.context,
		owner,
		name,
		nil,
	)

	return result, e
}

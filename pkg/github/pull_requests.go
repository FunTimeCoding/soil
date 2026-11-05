package github

import (
	"github.com/funtimecoding/soil/pkg/github/refusal"
	"github.com/google/go-github/v93/github"
)

func (c *Client) PullRequests(
	owner string,
	name string,
) ([]*github.PullRequest, error) {
	result, _, e := c.client.PullRequests.List(c.context, owner, name, nil)

	return result, refusal.Classify(e)
}

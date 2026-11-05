package github

import (
	"github.com/funtimecoding/soil/pkg/github/refusal"
	"github.com/google/go-github/v93/github"
)

func (c *Client) Organizations(user string) ([]*github.Organization, error) {
	result, _, e := c.client.Organizations.List(c.context, user, nil)

	return result, refusal.Classify(e)
}

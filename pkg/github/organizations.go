package github

import "github.com/google/go-github/v89/github"

func (c *Client) Organizations(user string) ([]*github.Organization, error) {
	result, _, e := c.client.Organizations.List(c.context, user, nil)

	return result, e
}

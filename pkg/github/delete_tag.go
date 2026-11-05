package github

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/github/refusal"
)

func (c *Client) DeleteTag(
	owner string,
	repository string,
	name string,
) error {
	_, e := c.client.Git.DeleteRef(
		c.context,
		owner,
		repository,
		fmt.Sprintf("refs/tags/%s", name),
	)

	return refusal.Classify(e)
}

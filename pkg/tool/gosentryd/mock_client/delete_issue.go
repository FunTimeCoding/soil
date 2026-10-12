package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"slices"
)

func (c *Client) DeleteIssue(
	_ string,
	identifier string,
) error {
	i := c.issueIndex(identifier)

	if i < 0 {
		return not_found.New("issue", identifier)
	}

	c.issues = slices.Delete(c.issues, i, i+1)

	return nil
}

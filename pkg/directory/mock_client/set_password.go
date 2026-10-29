package mock_client

import (
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) SetPassword(
	distinguishedName string,
	_ string,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	if _, exists := c.records[distinguishedName]; !exists {
		return not_found.New(constant.Subject, distinguishedName)
	}

	return nil
}

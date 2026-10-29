package mock_client

import (
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"slices"
)

func (c *Client) Delete(distinguishedName string) error {
	if c.Failure != nil {
		return c.Failure
	}

	if _, exists := c.records[distinguishedName]; !exists {
		return not_found.New(constant.Subject, distinguishedName)
	}

	delete(c.records, distinguishedName)
	c.order = slices.DeleteFunc(
		c.order,
		func(v string) bool { return v == distinguishedName },
	)

	return nil
}

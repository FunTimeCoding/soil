package mock_client

import (
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"maps"
)

func (c *Client) Add(
	distinguishedName string,
	attributes map[string][]string,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	if _, exists := c.records[distinguishedName]; exists {
		return conflict.Exists(constant.Subject, distinguishedName)
	}

	c.records[distinguishedName] = maps.Clone(attributes)
	c.order = append(c.order, distinguishedName)

	return nil
}

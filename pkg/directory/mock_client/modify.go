package mock_client

import (
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) Modify(
	distinguishedName string,
	attributes map[string][]string,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	current, exists := c.records[distinguishedName]

	if !exists {
		return not_found.New(constant.Subject, distinguishedName)
	}

	for name, values := range attributes {
		if len(values) == 0 {
			delete(current, name)

			continue
		}

		current[name] = values
	}

	return nil
}

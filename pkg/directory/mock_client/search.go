package mock_client

import (
	"github.com/funtimecoding/soil/pkg/directory/types/search_record"
	"maps"
	"slices"
	"strings"
)

func (c *Client) Search(
	filter string,
	_ []string,
) ([]*search_record.Record, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	class := strings.TrimSuffix(
		strings.TrimPrefix(filter, "(objectClass="),
		")",
	)
	var result []*search_record.Record

	for _, name := range c.order {
		attributes := c.records[name]

		if slices.Contains(attributes["objectClass"], class) {
			result = append(
				result,
				search_record.New(name, maps.Clone(attributes)),
			)
		}
	}

	return result, nil
}

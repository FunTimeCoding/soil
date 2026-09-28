package mock_client

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"strings"
)

func (c *Client) SearchItems(
	term string,
	types []string,
	page int,
	perPage int,
) ([]*item.Item, int, error) {
	var matched []*item.Item

	for _, i := range c.items {
		if term != "" && !strings.Contains(
			strings.ToLower(i.Name),
			strings.ToLower(term),
		) {
			continue
		}

		if len(types) > 0 && !containsType(types, i.Type) {
			continue
		}

		matched = append(matched, i)
	}

	total := len(matched)

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 25
	}

	start := (page - 1) * perPage

	if start >= total {
		return nil, total, nil
	}

	end := start + perPage

	if end > total {
		end = total
	}

	return matched[start:end], total, nil
}

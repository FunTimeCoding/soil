package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"slices"
	"strings"
)

func (c *Client) GetSearch(
	_ context.Context,
	p *client.GetSearchParams,
) (*client.SearchOutcome, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	collection := ""

	if p.Collection != nil {
		collection = *p.Collection
	}

	var found []client.SearchResult

	for _, d := range c.inCollection(collection) {
		if p.Exclude != nil && slices.Contains(*p.Exclude, d.Path) {
			continue
		}

		if !matches(d, p.Metadata) || d.Body == nil {
			continue
		}

		body := strings.ToLower(*d.Body)

		for _, word := range strings.Fields(strings.ToLower(p.Query)) {
			if strings.Contains(body, word) {
				d.Score = 1
				found = append(found, d)

				break
			}
		}
	}

	if p.Limit != nil {
		found = found[:min(*p.Limit, len(found))]
	}

	return &client.SearchOutcome{Results: found}, nil
}

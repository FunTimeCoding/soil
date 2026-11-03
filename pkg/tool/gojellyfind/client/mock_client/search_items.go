package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"slices"
	"strings"
)

func (c *Client) SearchItems(
	_ context.Context,
	p *client.SearchItemsParams,
) (*client.SearchResult, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	var types []string

	if p.Types != nil {
		for _, t := range strings.Split(*p.Types, ",") {
			types = append(types, strings.TrimSpace(t))
		}
	}

	var matched []client.Item

	for _, i := range c.ItemList {
		if p.Q != "" &&
			!strings.Contains(strings.ToLower(i.Name), strings.ToLower(p.Q)) {
			continue
		}

		if len(types) > 0 && !slices.Contains(types, i.Type) {
			continue
		}

		matched = append(matched, *i)
	}

	page, perPage := 1, 25

	if p.Page != nil {
		page = max(*p.Page, 1)
	}

	if p.PerPage != nil && *p.PerPage > 0 {
		perPage = *p.PerPage
	}

	result := &client.SearchResult{Total: len(matched), Page: page}
	start := (page - 1) * perPage

	if start < len(matched) {
		result.Items = matched[start:min(start+perPage, len(matched))]
	}

	return result, nil
}

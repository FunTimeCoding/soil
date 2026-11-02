package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"maps"
	"slices"
)

func (c *Client) ListSnippets(
	_ context.Context,
	_ *client.ListSnippetsParams,
) ([]client.SnippetItem, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	var result []client.SnippetItem

	for _, name := range slices.Sorted(maps.Keys(c.Snippets)) {
		result = append(result, client.SnippetItem{Name: name})
	}

	return result, nil
}

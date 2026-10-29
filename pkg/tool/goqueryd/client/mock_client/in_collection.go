package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"sort"
)

func (c *Client) inCollection(collection string) []client.SearchResult {
	var result []client.SearchResult

	for _, d := range c.documents {
		if collection == "" || d.Collection == collection {
			result = append(result, *d)
		}
	}

	sort.Slice(
		result,
		func(
			i int,
			j int,
		) bool {
			return result[i].Path < result[j].Path
		},
	)

	return result
}

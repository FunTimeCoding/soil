package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"slices"
)

func matches(
	d client.SearchResult,
	metadata *map[string]string,
) bool {
	if metadata == nil {
		return true
	}

	for k, v := range *metadata {
		if d.Metadata == nil || !slices.Contains((*d.Metadata)[k], v) {
			return false
		}
	}

	return true
}

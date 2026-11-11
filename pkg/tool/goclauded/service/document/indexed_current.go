package document

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"slices"
)

func indexedCurrent(
	indexed map[string]client.SearchResult,
	path string,
	hash string,
	metadata map[string][]string,
) bool {
	r, found := indexed[path]

	if !found || r.Hash != hash {
		return false
	}

	if r.Metadata == nil {
		return len(metadata) == 0
	}

	for key, values := range metadata {
		if !slices.Equal((*r.Metadata)[key], values) {
			return false
		}
	}

	return true
}

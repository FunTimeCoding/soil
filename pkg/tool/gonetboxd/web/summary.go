package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/bookmark"
)

func summary(bookmarks []*bookmark.Bookmark) []string {
	if len(bookmarks) == 0 {
		return nil
	}

	kinds := map[string]bool{}

	for _, b := range bookmarks {
		kinds[b.ObjectType] = true
	}

	result := []string{fmt.Sprintf("%d bookmarked", len(bookmarks))}

	if len(kinds) > 1 {
		result = append(result, fmt.Sprintf("%d types", len(kinds)))
	}

	return result
}

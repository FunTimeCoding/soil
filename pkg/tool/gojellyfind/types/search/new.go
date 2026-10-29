package search

import "github.com/funtimecoding/soil/pkg/jellyfin/item"

func New(
	items []*item.Item,
	total int,
	page int,
) *Result {
	return &Result{Total: total, Page: page, Items: items}
}

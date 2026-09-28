package server

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func convertItems(items []*item.Item) []*server.Item {
	result := make([]*server.Item, 0, len(items))

	for _, v := range items {
		result = append(result, convertItem(v))
	}

	return result
}

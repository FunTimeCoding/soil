package search

import "github.com/funtimecoding/soil/pkg/jellyfin/item"

type Result struct {
	Total int          `json:"total"`
	Page  int          `json:"page"`
	Items []*item.Item `json:"items"`
}

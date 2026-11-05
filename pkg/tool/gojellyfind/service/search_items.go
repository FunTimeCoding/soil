package service

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/types/search"
)

func (s *Service) SearchItems(
	term string,
	types []string,
	page int,
	perPage int,
) (*search.Result, error) {
	p := max(page, 1)
	var items []*item.Item
	var total int
	e := s.attempt(
		func() error {
			var f error
			items, total, f = s.client.SearchItems(term, types, p, perPage)

			return f
		},
	)

	if e != nil {
		return nil, e
	}

	return search.New(items, total, p), nil
}

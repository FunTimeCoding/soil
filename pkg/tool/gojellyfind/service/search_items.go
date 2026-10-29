package service

import "github.com/funtimecoding/soil/pkg/tool/gojellyfind/types/search"

func (s *Service) SearchItems(
	term string,
	types []string,
	page int,
	perPage int,
) (*search.Result, error) {
	p := max(page, 1)
	items, total, e := s.source.SearchItems(term, types, p, perPage)

	if e != nil {
		return nil, e
	}

	return search.New(items, total, p), nil
}
